package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/currentdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/daydata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/logindata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/monthdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/salesforcedata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/yeardata"
	"golang.org/x/net/html"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/rs/zerolog/log"
)

func NewSolarZeroScrape(config *config.Configuration) *SolarZeroScrape {
	s := &SolarZeroScrape{
		config:         config,
		userAttributes: make(map[string]string),
		salesForceData: salesforcedata.SalesForceData{},
		reauthenticate: false,
	}

	return s
}

type SolarZeroScrape struct {
	config         *config.Configuration
	accessToken    string
	refreshToken   string
	idToken        string
	cognitoSvc     *cip.Client
	userAttributes map[string]string
	salesForceData salesforcedata.SalesForceData
	cookies        []*http.Cookie
	reauthenticate bool

	logindata logindata.LoginData

	currentData currentdata.CurrentData
	dayData     daydata.DayData
	monthData   monthdata.MonthData
	yearData    yeardata.YearData
}

func (szs *SolarZeroScrape) cognitoAuth() bool {
	log.Info().Msg("INFO: Starting Cognito Authentication")
	csrp, _ := cognitosrp.NewCognitoSRP(
		szs.config.SolarZero.Username,
		szs.config.SolarZero.Password,
		szs.config.SolarZero.UserPoolID,
		szs.config.SolarZero.ClientID, nil)

	// configure cognito identity provider
	cfg, _ := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(szs.config.SolarZero.API.Region),
		awsconfig.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)
	svc := cip.NewFromConfig(cfg)

	// initiate auth
	resp, err := svc.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: csrp.GetAuthParams(),
	})
	if err != nil {
		log.Error().Msgf("Cognito Authentication (InitiateAuth): %s ", err.Error())
		return false
	}

	// respond to password verifier challenge
	if resp.ChallengeName == types.ChallengeNameTypePasswordVerifier {
		challengeResponses, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())

		resp, err := svc.RespondToAuthChallenge(context.Background(), &cip.RespondToAuthChallengeInput{
			ChallengeName:      types.ChallengeNameTypePasswordVerifier,
			ChallengeResponses: challengeResponses,
			ClientId:           aws.String(csrp.GetClientId()),
		})
		if err != nil {
			log.Error().Msgf("Cognito Authentication (RespondToAuthChallenge): %s", err.Error())
			return false
		}

		szs.accessToken = *resp.AuthenticationResult.AccessToken
		szs.idToken = *resp.AuthenticationResult.IdToken
		szs.refreshToken = *resp.AuthenticationResult.RefreshToken
		szs.cognitoSvc = svc
		// print the tokens
		log.Debug().Str("Access Token", *resp.AuthenticationResult.AccessToken)
		log.Debug().Str("ID Token", *resp.AuthenticationResult.IdToken)
		log.Debug().Str("Refresh Token", *resp.AuthenticationResult.RefreshToken)
		log.Info().Msg("Cognito Authentication Success")
		return szs.getUser()
	} else {
		log.Error().Msgf("Cognito Authentication (Unhandled Challenge): %s", resp.ChallengeName)
		return false
	}
}

func (szs *SolarZeroScrape) getUser() bool {
	log.Info().Msg("Cognito GetUser")

	getUserOutput, err := szs.cognitoSvc.GetUser(context.Background(), &cip.GetUserInput{
		AccessToken: &szs.accessToken,
	})

	if err != nil {
		log.Error().Msgf("Cognito GetUser (GetUser): %s", err.Error())
		return false
	}
	for _, attr := range getUserOutput.UserAttributes {
		szs.userAttributes[*attr.Name] = *attr.Value
	}

	log.Info().Msg("Cognito GetUser Success")

	return true
}

func (szs *SolarZeroScrape) fetchSalesForceData() bool {
	client := &http.Client{}

	log.Info().Msg("Fetch SalesForce Data")

	var jsonStr = []byte(`"` + szs.userAttributes["custom:contactId"] + `"`)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/prod/newuserinfo",
		szs.config.SolarZero.API.APIGatewayURL), bytes.NewBuffer(jsonStr))

	req.Header.Add("X-API-KEY", szs.config.SolarZero.API.APIKey)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Fetch SalesForce Data (Request): %s", err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("Fetch SalesForce Data (ReadAll): %s", err.Error())
		return false
	}

	log.Debug().RawJSON("SalesForceData", body)

	szs.salesForceData, err = salesforcedata.UnmarshalSalesForceData(body)
	if err != nil {
		log.Error().Msgf("Fetch SalesForce Data (Unmarshal): %s", err.Error())
		return false
	}
	log.Info().Msg("Fetch SalesForce Data Success")

	return true
}

func (szs *SolarZeroScrape) getCookies() bool {
	szs.reauthenticate = false

	log.Info().Msg("Get Cookies and Login Data")

	url := fmt.Sprintf("https://%s/login/%s",
		szs.config.SolarZero.API.SolarZeroAPIAddress, szs.salesForceData.Token)
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Get Cookies and Login Data (Login Request): %s", err.Error())
		return false
	}

	cookies := res.Cookies()
	location := res.Header.Get("Location")
	res.Body.Close()

	if cookies == nil {
		log.Error().Msgf("Get Cookies and Login Data (No Cookies 1): %s", err.Error())
		return false
	}

	client = &http.Client{}
	url = fmt.Sprintf("https://%s%s",
		szs.config.SolarZero.API.SolarZeroAPIAddress, location)

	// 2nd stage of cookie get.
	req, _ = http.NewRequest(method, url, nil)

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	res, err = client.Do(req)
	if err != nil {
		log.Error().Msgf("Get Cookies and Login Data (Data Request): %s", err.Error())
		return false
	}
	defer res.Body.Close()
	// print(res.Cookies())

	szs.cookies = res.Cookies()
	if szs.cookies == nil {
		log.Error().Msgf("Get Cookies and Login Data (No Cookies 2): %s", err.Error())
		return false
	}

	z := html.NewTokenizer(res.Body)
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			log.Error().Msgf("Get Cookies and Login Data (ErrorToken): " + z.Err().Error())
			return false
		case html.TextToken:
			if depth > 0 {
				text := string(z.Text())
				if strings.HasPrefix(text, "window.__data__ = ") {
					text = strings.Replace(text, "window.__data__ = ", "", 1)
					szs.logindata, err = logindata.UnmarshalLoginData([]byte(text))
					if err != nil {
						log.Error().Msgf("Get Cookies and Login Data (UnmarshalLoginData): %s", err.Error())
						return false
					}
					log.Debug().RawJSON("LoginData", []byte(text))

					log.Info().Msg("Get Cookies and Login Data Success")
					return true
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if string(tn) == "script" {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}

}

func (szs *SolarZeroScrape) getWithCookies(url string) ([]byte, error) {
	log.Debug().Msg("Get Url With Cookies: " + url)

	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest(method, url, nil)

	for _, cookie := range szs.cookies {
		req.AddCookie(cookie)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Error().Msgf("Get Url With Cookies (Request): %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Error().Msgf("Get Url With Cookies (Status Code): " + res.Status)
		szs.reauthenticate = true
		return nil, fmt.Errorf("needs reauthentication")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msgf("Get Url With Cookies (ReadAll): %s", err.Error())
		return nil, err
	}

	log.Debug().Msg("Get Url With Cookies Success")

	return body, nil

}

func (szs *SolarZeroScrape) getCurrentData() bool {
	log.Info().Msg("Get Current Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getCurrentData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		log.Error().Msgf("Get Current Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.currentData, err = currentdata.UnmarshalCurrentData(body)
	if err != nil {
		log.Error().Msgf("Get Current Data (UnmarshalCurrentData): %s", err.Error())
		return false
	}
	log.Debug().RawJSON("CurrentData", body)
	log.Info().Msg("Get Current Data Success")
	return true
}

func (szs *SolarZeroScrape) getDayData() bool {
	log.Info().Msg("Get Day Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getDayData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		log.Error().Msgf("Get Day Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.dayData, err = daydata.UnmarshalDayData(body)
	if err != nil {
		log.Error().Msgf("Get Day Data (UnmarshalDayData): %s", err.Error())
		return false
	}
	log.Debug().RawJSON("DayData", body)
	log.Info().Msg("Get Day Data Success")
	return true
}

func (szs *SolarZeroScrape) getMonthData() bool {
	log.Info().Msg("Get Month Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getMonthData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		log.Error().Msgf("Get Month Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.monthData, err = monthdata.UnmarshalMonthData(body)
	if err != nil {
		log.Error().Msgf("Get Month Data (UnmarshalMonthData): %s", err.Error())
		return false
	}
	log.Debug().RawJSON("MonthData", body)
	log.Info().Msg("Get Month Data Success")
	return true
}

func (szs *SolarZeroScrape) getYearData() bool {
	log.Info().Msg("Get Year Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getYearData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		log.Error().Msgf("Get Year Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.yearData, err = yeardata.UnmarshalYearData(body)
	if err != nil {
		log.Error().Msgf("Get Year Data (UnmarshalYearData): %s", err.Error())
		return false
	}
	log.Debug().RawJSON("YearData", body)
	log.Info().Msg("Get Year Data Success")
	return true
}

func (szs *SolarZeroScrape) AuthenticateFully() bool {
	if !szs.cognitoAuth() {
		return false
	}

	if !szs.fetchSalesForceData() {
		return false
	}

	return szs.getCookies()
}

func (szs *SolarZeroScrape) GetData() bool {
	success := szs.getCurrentData()
	if !success {
		if szs.reauthenticate {
			if szs.getCookies() {
				success = szs.getCurrentData()
				if !success {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if !szs.getDayData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getDayData() {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if !szs.getMonthData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getMonthData() {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if !szs.getYearData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getYearData() {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
