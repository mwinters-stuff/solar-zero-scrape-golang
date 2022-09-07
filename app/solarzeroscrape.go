package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/currentdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/daydata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/logindata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/monthdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/yeardata"
	"golang.org/x/net/html"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func NewSolarZeroScrape(config *config.Configuration) *SolarZeroScrape {
	s := &SolarZeroScrape{
		config:         config,
		userAttributes: make(map[string]string),
		salesForceData: make(map[string]interface{}),
		reauthenticate: false,
	}

	return s
}

type SolarZeroScrape struct {
	config          *config.Configuration
	accessToken     string
	refreshToken    string
	idToken         string
	cognitoSvc      *cip.Client
	userAttributes  map[string]string
	salesForceData  map[string]interface{}
	salesForceToken string
	cookies         []*http.Cookie
	reauthenticate  bool

	logindata logindata.LoginData

	currentData currentdata.CurrentData
	dayData     daydata.DayData
	monthData   monthdata.MonthData
	yearData    yeardata.YearData
}

func (szs *SolarZeroScrape) cognitoAuth() bool {
	println("INFO: Starting Cognito Authentication")
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
		println("ERROR: Cognito Authentication (InitiateAuth): " + err.Error())
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
			println("ERROR: Cognito Authentication (RespondToAuthChallenge): " + err.Error())
			return false
		}

		szs.accessToken = *resp.AuthenticationResult.AccessToken
		szs.idToken = *resp.AuthenticationResult.IdToken
		szs.refreshToken = *resp.AuthenticationResult.RefreshToken
		szs.cognitoSvc = svc
		// print the tokens
		// fmt.Printf("Access Token: %s\n", *resp.AuthenticationResult.AccessToken)
		// fmt.Printf("ID Token: %s\n", *resp.AuthenticationResult.IdToken)
		// fmt.Printf("Refresh Token: %s\n", *resp.AuthenticationResult.RefreshToken)
		println("INFO: Cognito Authentication Success")
		return szs.getUser()
	} else {
		println("ERROR: Cognito Authentication (Unhandled Challenge): " + resp.ChallengeName)
		return false
	}
}

func (szs *SolarZeroScrape) getUser() bool {
	println("INFO: Cognito GetUser")

	getUserOutput, err := szs.cognitoSvc.GetUser(context.Background(), &cip.GetUserInput{
		AccessToken: &szs.accessToken,
	})

	if err != nil {
		println("ERROR: Cognito GetUser (GetUser): " + err.Error())
		return false
	}
	for i := 0; i < len(getUserOutput.UserAttributes); i++ {
		var attr = getUserOutput.UserAttributes[i]
		szs.userAttributes[*attr.Name] = *attr.Value
	}

	println("INFO: Cognito GetUser Success")

	return true
}

func (szs *SolarZeroScrape) fetchSalesForceData() bool {
	client := &http.Client{}

	println("INFO: Fetch SalesForce Data")

	var jsonStr = []byte(`"` + szs.userAttributes["custom:contactId"] + `"`)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/prod/newuserinfo",
		szs.config.SolarZero.API.APIGatewayURL), bytes.NewBuffer(jsonStr))
	// ...
	req.Header.Add("X-API-KEY", szs.config.SolarZero.API.APIKey)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		println("ERROR: Fetch SalesForce Data (Request): " + err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("ERROR: Fetch SalesForce Data (ReadAll): " + err.Error())
		return false
	}
	// fmt.Println(string(body))

	// var dat map[string]interface{}
	if err := json.Unmarshal(body, &szs.salesForceData); err != nil {
		println("ERROR: Fetch SalesForce Data (Unmarshal): " + err.Error())
		return false
	}
	szs.salesForceToken = szs.salesForceData["token"].(string)
	// fmt.Printf("SalesForceToken: %s\n", szs.salesForceToken)
	println("INFO: Fetch SalesForce Data Success")

	return true
}

func (szs *SolarZeroScrape) getCookies() bool {

	println("INFO: Get Cookies and Login Data")

	url := fmt.Sprintf("https://%s/login/%s",
		szs.config.SolarZero.API.SolarZeroAPIAddress, szs.salesForceToken)
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		println("ERROR: Get Cookies and Login Data (Login Request): " + err.Error())
		return false
	}

	cookies := res.Cookies()
	location := res.Header.Get("Location")
	res.Body.Close()

	if cookies == nil {
		println("ERROR: Get Cookies and Login Data (No Cookies 1): " + err.Error())
		return false
	}

	client = &http.Client{}
	url = fmt.Sprintf("https://%s%s",
		szs.config.SolarZero.API.SolarZeroAPIAddress, location)

	// 2nd stage of cookie get.
	req, _ = http.NewRequest(method, url, nil)

	for i := 0; i < len(cookies); i++ {
		req.AddCookie(cookies[i])
	}

	res, err = client.Do(req)
	if err != nil {
		println("ERROR: Get Cookies and Login Data (Data Request): " + err.Error())
		return false
	}
	defer res.Body.Close()
	// print(res.Cookies())

	szs.cookies = res.Cookies()
	if szs.cookies == nil {
		println("ERROR: Get Cookies and Login Data (No Cookies 2): " + err.Error())
		return false
	}

	z := html.NewTokenizer(res.Body)
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			println("ERROR: Get Cookies and Login Data (ErrorToken): " + z.Err().Error())
			return false
		case html.TextToken:
			if depth > 0 {
				text := string(z.Text())
				if strings.HasPrefix(text, "window.__data__ = ") {
					text = strings.Replace(text, "window.__data__ = ", "", 1)
					szs.logindata, err = logindata.UnmarshalLoginData([]byte(text))
					if err != nil {
						println("ERROR: Get Cookies and Login Data (UnmarshalLoginData): " + err.Error())
						return false
					}

					println("INFO: Get Cookies and Login Data Success")
					szs.reauthenticate = false
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
	println("INFO: Get Url With Cookies: " + url)

	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest(method, url, nil)

	for i := 0; i < len(szs.cookies); i++ {
		req.AddCookie(szs.cookies[i])
	}

	res, err := client.Do(req)
	if err != nil {
		println("ERROR: Get Url With Cookies (Request): " + err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		println("ERROR: Get Url With Cookies (Status Code): " + res.Status)
		szs.reauthenticate = true
		return nil, fmt.Errorf("needs reauthentication")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println("ERROR: Get Url With Cookies (ReadAll): " + err.Error())
		return nil, err
	}
	// fmt.Println(string(body))
	println("INFO: Get Url With Cookies Success")

	return body, nil

}

func (szs *SolarZeroScrape) getCurrentData() (changed bool, success bool) {
	println("INFO: Get Current Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getCurrentData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		println("ERROR: Get Current Data (getWithCookies): " + err.Error())
		return false, false
	}
	lastData := szs.currentData
	szs.currentData, err = currentdata.UnmarshalCurrentData(body)
	if err != nil {
		println("ERROR: Get Current Data (UnmarshalCurrentData): " + err.Error())
		return false, false
	}
	// fmt.Printf("Got Current Data: %s\n", string(body))
	println("INFO: Get Current Data Success")
	return !lastData.Equals(&szs.currentData), true
}

func (szs *SolarZeroScrape) getDayData() bool {
	println("INFO: Get Day Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getDayData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		println("ERROR: Get Day Data (getWithCookies): " + err.Error())
		return false
	}
	szs.dayData, err = daydata.UnmarshalDayData(body)
	if err != nil {
		println("ERROR: Get Day Data (UnmarshalDayData): " + err.Error())
		return false
	}
	// fmt.Printf("Got Day Data: %s\n", string(body))
	println("INFO: Get Day Data Success")
	return true
}

func (szs *SolarZeroScrape) getMonthData() bool {
	println("INFO: Get Month Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getMonthData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		println("ERROR: Get Month Data (getWithCookies): " + err.Error())
		return false
	}
	szs.monthData, err = monthdata.UnmarshalMonthData(body)
	if err != nil {
		println("ERROR: Get Month Data (UnmarshalMonthData): " + err.Error())
		return false
	}
	// fmt.Printf("Got Month Data: %s\n", string(body))
	println("INFO: Get Month Data Success")
	return true
}

func (szs *SolarZeroScrape) getYearData() bool {
	println("INFO: Get Year Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getYearData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		println("ERROR: Get Year Data (getWithCookies): " + err.Error())
		return false
	}
	szs.yearData, err = yeardata.UnmarshalYearData(body)
	if err != nil {
		println("ERROR: Get Year Data (UnmarshalYearData): " + err.Error())
		return false
	}
	// fmt.Printf("Got Year Data: %s\n", string(body))
	println("INFO: Get Year Data Success")
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

func (szs *SolarZeroScrape) GetData() (bool, bool) {
	changed, success := szs.getCurrentData()
	if !success {
		if szs.reauthenticate {
			if szs.getCookies() {
				changed, success = szs.getCurrentData()
				if !success {
					return false, false
				}
			} else {
				return false, false
			}
		} else {
			return false, false
		}
	}

	if !szs.getDayData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getDayData() {
					return false, false
				}
			} else {
				return false, false
			}
		} else {
			return false, false
		}
	}

	if !szs.getMonthData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getMonthData() {
					return false, false
				}
			} else {
				return false, false
			}
		} else {
			return false, false
		}
	}

	if !szs.getYearData() {
		if szs.reauthenticate {
			if szs.getCookies() {
				if !szs.getYearData() {
					return false, false
				}
			} else {
				return false, false
			}
		} else {
			return false, false
		}
	}

	return changed, true
}
