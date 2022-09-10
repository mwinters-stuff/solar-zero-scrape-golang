package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, attr := range getUserOutput.UserAttributes {
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

	req.Header.Add("X-API-KEY", szs.config.SolarZero.API.APIKey)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		println("ERROR: Fetch SalesForce Data (Request): " + err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("ERROR: Fetch SalesForce Data (ReadAll): " + err.Error())
		return false
	}
	// fmt.Println(string(body))
	szs.writeToLog("SalesForceData", body)

	szs.salesForceData, err = salesforcedata.UnmarshalSalesForceData(body)
	if err != nil {
		println("ERROR: Fetch SalesForce Data (Unmarshal): " + err.Error())
		return false
	}
	println("INFO: Fetch SalesForce Data Success")

	return true
}

func (szs *SolarZeroScrape) getCookies() bool {

	println("INFO: Get Cookies and Login Data")

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

	for _, cookie := range cookies {
		req.AddCookie(cookie)
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
					szs.writeToLog("LoginData", []byte(text))

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

	for _, cookie := range szs.cookies {
		req.AddCookie(cookie)
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		println("ERROR: Get Url With Cookies (ReadAll): " + err.Error())
		return nil, err
	}
	szs.writeToLog(url, body)
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

func (szs *SolarZeroScrape) writeToLog(what string, data []byte) {
	if szs.config.DebugLog != nil {

		file, err := os.OpenFile(*szs.config.DebugLog, os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			defer file.Close()
			file.WriteString(what)
			file.WriteString("\n")
			file.Write(data)
			file.WriteString("\n\n\n")
		}
	}
}

//https://solarzero.pnz.technology/api/getRangeData/data?id=18003&api=https://ems.solarcity.panabattery.com/api&from=2022-07-01&to=2022-08-06
//[{"Date":"1-Jul","Solar use":6.029999999999999,"Grid":42.71,"Export":-0.07,"Load":48.34,"Solar":6.1,"Discharge":9.4,"Charge":9.8},{"Date":"2-Jul","Solar use":5.63,"Grid":40.06,"Export":-0.07,"Load":45.29,"Solar":5.7,"Discharge":9.4,"Charge":9.8},{"Date":"3-Jul","Solar use":6.05,"Grid":39.58,"Export":-0.05,"Load":45.23,"Solar":6.1,"Discharge":9.4,"Charge":9.8},{"Date":"4-Jul","Solar use":5.78,"Grid":34.19,"Export":-0.12,"Load":39.67,"Solar":5.9,"Discharge":9.5,"Charge":9.8},{"Date":"5-Jul","Solar use":1.89,"Grid":39.69,"Export":-0.01,"Load":41.08,"Solar":1.9,"Discharge":9.5,"Charge":10},{"Date":"6-Jul","Solar use":6.029999999999999,"Grid":37.2,"Export":-0.07,"Load":42.83,"Solar":6.1,"Discharge":9.8,"Charge":10.2},{"Date":"7-Jul","Solar use":2.5700000000000003,"Grid":48.47,"Export":-0.03,"Load":50.54,"Solar":2.6,"Discharge":9.7,"Charge":10.2},{"Date":"8-Jul","Solar use":1.29,"Grid":42.72,"Export":-0.01,"Load":43.61,"Solar":1.3,"Discharge":9.6,"Charge":10},{"Date":"9-Jul","Solar use":0.62,"Grid":30.9,"Export":-0.08,"Load":31.02,"Solar":0.7,"Discharge":9.5,"Charge":10},{"Date":"10-Jul","Solar use":4.16,"Grid":26.68,"Export":-0.14,"Load":30.34,"Solar":4.3,"Discharge":8.4,"Charge":8.9},{"Date":"11-Jul","Solar use":1.89,"Grid":59.34,"Export":-0.01,"Load":60.73,"Solar":1.9,"Discharge":9.4,"Charge":9.9},{"Date":"12-Jul","Solar use":0,"Grid":51.62,"Export":-0.03,"Load":51.2,"Solar":0,"Discharge":9.7,"Charge":10.1},{"Date":"13-Jul","Solar use":3.77,"Grid":32.29,"Export":-0.03,"Load":35.56,"Solar":3.8,"Discharge":9.5,"Charge":10},{"Date":"14-Jul","Solar use":5.9,"Grid":23.23,"Export":-0.1,"Load":28.73,"Solar":6,"Discharge":9.5,"Charge":9.9},{"Date":"15-Jul","Solar use":3.68,"Grid":42.5,"Export":-0.02,"Load":45.78,"Solar":3.7,"Discharge":9.5,"Charge":9.9},{"Date":"16-Jul","Solar use":6.46,"Grid":35.39,"Export":-0.04,"Load":41.35,"Solar":6.5,"Discharge":9.5,"Charge":10},{"Date":"17-Jul","Solar use":1.79,"Grid":53.36,"Export":-0.01,"Load":54.65,"Solar":1.8,"Discharge":9.7,"Charge":10.2},{"Date":"18-Jul","Solar use":4.43,"Grid":28.64,"Export":-0.07,"Load":32.77,"Solar":4.5,"Discharge":10.2,"Charge":10.5},{"Date":"19-Jul","Solar use":1.9500000000000002,"Grid":37.63,"Export":-0.15,"Load":39.18,"Solar":2.1,"Discharge":9.6,"Charge":10},{"Date":"20-Jul","Solar use":1.09,"Grid":45.93,"Export":-0.01,"Load":46.51,"Solar":1.1,"Discharge":9.5,"Charge":10},{"Date":"21-Jul","Solar use":3.29,"Grid":51.17,"Export":-0.01,"Load":53.96,"Solar":3.3,"Discharge":9.4,"Charge":9.9},{"Date":"22-Jul","Solar use":6.52,"Grid":42.34,"Export":-0.08,"Load":48.36,"Solar":6.6,"Discharge":10.1,"Charge":10.6},{"Date":"23-Jul","Solar use":6.46,"Grid":33.77,"Export":-0.14,"Load":39.93,"Solar":6.6,"Discharge":9.6,"Charge":9.9},{"Date":"24-Jul","Solar use":3.0300000000000002,"Grid":39.54,"Export":-0.07,"Load":42.08,"Solar":3.1,"Discharge":9.5,"Charge":10},{"Date":"25-Jul","Solar use":1.68,"Grid":40.06,"Export":-0.02,"Load":41.23,"Solar":1.7,"Discharge":9.7,"Charge":10.2},{"Date":"26-Jul","Solar use":0.38,"Grid":43.12,"Export":-0.02,"Load":43.1,"Solar":0.4,"Discharge":9.5,"Charge":9.9},{"Date":"27-Jul","Solar use":5.46,"Grid":34.79,"Export":-0.04,"Load":39.74,"Solar":5.5,"Discharge":9.5,"Charge":10},{"Date":"28-Jul","Solar use":1.15,"Grid":40.36,"Export":-0.05,"Load":41.1,"Solar":1.2,"Discharge":9.5,"Charge":9.9},{"Date":"29-Jul","Solar use":6.46,"Grid":39.8,"Export":-0.04,"Load":45.75,"Solar":6.5,"Discharge":9.5,"Charge":10},{"Date":"30-Jul","Solar use":1.39,"Grid":47.02,"Export":-0.01,"Load":47.91,"Solar":1.4,"Discharge":9.5,"Charge":10},{"Date":"31-Jul","Solar use":6.71,"Grid":42.39,"Export":-0.09,"Load":48.6,"Solar":6.8,"Discharge":9.4,"Charge":9.9},{"Date":"1-Aug","Solar use":7.88,"Grid":34.45,"Export":-0.12,"Load":41.73,"Solar":8,"Discharge":6.7,"Charge":7.3},{"Date":"2-Aug","Solar use":7.31,"Grid":22.08,"Export":-0.19,"Load":29.3,"Solar":7.5,"Discharge":8.9,"Charge":9},{"Date":"3-Aug","Solar use":7.6000000000000005,"Grid":14.37,"Export":-0.3,"Load":21.57,"Solar":7.9,"Discharge":8.5,"Charge":8.9},{"Date":"4-Aug","Solar use":6.6499999999999995,"Grid":20.23,"Export":-0.15,"Load":26.07,"Solar":6.8,"Discharge":8.7,"Charge":9.5},{"Date":"5-Aug","Solar use":2.87,"Grid":31.24,"Export":-0.03,"Load":33.51,"Solar":2.9,"Discharge":8.4,"Charge":9},{"Date":"6-Aug","Solar use":2.26,"Grid":28.31,"Export":-0.14,"Load":30.37,"Solar":2.4,"Discharge":8.6,"Charge":8.8}]
