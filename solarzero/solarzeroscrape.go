package solarzero

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"golang.org/x/net/html"
)

type SolarZeroScrape interface {
	Start()
	AuthenticateFully() bool
	GetData() bool
	CurrentData() jsontypes.CurrentData
	DayData() jsontypes.DayData
	MonthData() jsontypes.MonthData
	YearData() jsontypes.YearData

	Ready() bool
	Healthy() bool
}

type SolarZeroScrapeImpl struct {
	config jsontypes.Configuration
	// accessToken  string
	// refreshToken string
	// idToken      string

	// userAttributes map[string]string
	salesForceData jsontypes.SalesForceData
	cookies        []*http.Cookie
	reauthenticate bool

	influxdb InfluxDBWriter
	mqtt     MQTTClient

	logindata jsontypes.LoginData

	currentData jsontypes.CurrentData
	dayData     jsontypes.DayData
	monthData   jsontypes.MonthData
	yearData    jsontypes.YearData

	awsInterface AWSInterface

	lastGoodWriteTimestamp time.Time
	ready                  bool
}

func NewSolarZeroScrape(options *AllSolarZeroOptions) SolarZeroScrape {
	config := jsontypes.Configuration{}

	if options.SolarZeroOptions.Config != "" {
		var err error
		config, err = jsontypes.LoadConfiguration(options.SolarZeroOptions.Config)
		if err != nil {
			Logger.Panic().Msg("LoadConfiguration " + err.Error())
		}
	} else {
		config.SolarZero.Username = options.SolarZeroOptions.Username
		config.SolarZero.Password = options.SolarZeroOptions.Password
		config.SolarZero.UserPoolID = options.OtherOptions.UserPoolId
		config.SolarZero.ClientID = options.OtherOptions.ClientId
		config.SolarZero.API.Region = options.OtherOptions.ApiRegion
		config.SolarZero.API.APIKey = options.OtherOptions.ApiKey
		config.SolarZero.API.APIGatewayURL = options.OtherOptions.ApiGatewayURL
		config.SolarZero.API.SolarZeroAPIAddress = options.OtherOptions.ApiSolarZeroApiAddress

		config.InfluxDB.HostURL = options.InfluxDBOptions.HostURL
		config.InfluxDB.Token = options.InfluxDBOptions.Token
		config.InfluxDB.Org = options.InfluxDBOptions.Org
		config.InfluxDB.Bucket = options.InfluxDBOptions.Bucket

		config.Mqtt.URL = options.MQTTOptions.ServerURL
		config.Mqtt.Username = options.MQTTOptions.Username
		config.Mqtt.Password = options.MQTTOptions.Password
		config.Mqtt.BaseTopic = options.MQTTOptions.Topic

	}

	influxdb := NewInfluxDBWriter(&config)
	err := influxdb.Connect(influxdb2.NewClient(config.InfluxDB.HostURL, config.InfluxDB.Token))
	if err != nil {
		Logger.Panic().Msgf("InfluxDB Connect %s", err.Error())
	}

	var mqtt MQTTClient
	if config.Mqtt.URL != "" {
		mqtt = NewMQTTClient(&config)
		err = mqtt.Connect()
		if err != nil {
			Logger.Panic().Msgf("MQTT Connect %s", err.Error())
		}
	}

	Logger.Info().Msg("Authenticating")

	scrape := &SolarZeroScrapeImpl{
		awsInterface: NewAWSInterface(&config),
		config:       config,
		influxdb:     influxdb,
		mqtt:         mqtt,
		// userAttributes: make(map[string]string),
		salesForceData:         jsontypes.SalesForceData{},
		reauthenticate:         false,
		lastGoodWriteTimestamp: time.Now(),
		ready:                  false,
	}

	return scrape
}

func (szs *SolarZeroScrapeImpl) Start() {
	szs.ready = true

	s := gocron.NewScheduler(time.Local)

	for szs.AuthenticateFully() {
		s.Every(5).Minutes().Do(func() {
			Logger.Info().Msgf("Get Data at %s", time.Now())
			success := szs.GetData()
			if success {
				szs.influxdb.WriteData(szs)
				szs.lastGoodWriteTimestamp = time.Now()
				if szs.mqtt != nil {
					szs.mqtt.WriteData(szs)
				}
			} else {
				Logger.Error().Msg("GetData Failed, Reauthenticating")
				s.Stop()
			}
		})
		s.StartBlocking()
	}
	Logger.Error().Msg("AuthenicateFully Failed, Exiting")
}

func (szs *SolarZeroScrapeImpl) cognitoAuth() bool {
	if szs.awsInterface.Authenticate() {
		return szs.awsInterface.GetUser()
	} else {
		return false
	}
}

func (szs *SolarZeroScrapeImpl) fetchSalesForceData() bool {
	client := &http.Client{}

	Logger.Info().Msg("Fetch SalesForce Data")

	var jsonStr = []byte(`"` + szs.awsInterface.UserAttributes()["custom:contactId"] + `"`)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/prod/newuserinfo",
		szs.config.SolarZero.API.APIGatewayURL), bytes.NewBuffer(jsonStr))

	req.Header.Add("X-API-KEY", szs.config.SolarZero.API.APIKey)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		Logger.Error().Msgf("Fetch SalesForce Data (Request): %s", err.Error())
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Error().Msgf("Fetch SalesForce Data (ReadAll): %s", err.Error())
		return false
	}

	Logger.Debug().Msgf("SalesForceData: %s", body)

	szs.salesForceData, err = jsontypes.UnmarshalSalesForceData(body)
	if err != nil {
		Logger.Error().Msgf("Fetch SalesForce Data (Unmarshal): %s", err.Error())
		return false
	}
	Logger.Info().Msg("Fetch SalesForce Data Success")

	return true
}

func (szs *SolarZeroScrapeImpl) getCookies() bool {
	szs.reauthenticate = false

	Logger.Info().Msg("Get Cookies and Login Data")

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
		Logger.Error().Msgf("Get Cookies and Login Data (Login Request): %s", err.Error())
		return false
	}

	cookies := res.Cookies()
	location := res.Header.Get("Location")
	res.Body.Close()

	if cookies == nil {
		Logger.Error().Msgf("Get Cookies and Login Data (No Cookies 1): %s", err.Error())
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
		Logger.Error().Msgf("Get Cookies and Login Data (Data Request): %s", err.Error())
		return false
	}
	defer res.Body.Close()
	// print(res.Cookies())

	szs.cookies = res.Cookies()
	if szs.cookies == nil {
		Logger.Error().Msgf("Get Cookies and Login Data (No Cookies 2): %s", err.Error())
		return false
	}

	z := html.NewTokenizer(res.Body)
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			Logger.Error().Msgf("Get Cookies and Login Data (ErrorToken): " + z.Err().Error())
			return false
		case html.TextToken:
			if depth > 0 {
				text := string(z.Text())
				if strings.HasPrefix(text, "window.__data__ = ") {
					text = strings.Replace(text, "window.__data__ = ", "", 1)
					szs.logindata, err = jsontypes.UnmarshalLoginData([]byte(text))
					if err != nil {
						Logger.Error().Msgf("Get Cookies and Login Data (UnmarshalLoginData): %s", err.Error())
						return false
					}
					Logger.Debug().Msgf("LoginData %s", text)

					Logger.Info().Msg("Get Cookies and Login Data Success")
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

func (szs *SolarZeroScrapeImpl) getWithCookies(url string) ([]byte, error) {
	Logger.Debug().Msg("Get Url With Cookies: " + url)

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
		Logger.Error().Msgf("Get Url With Cookies (Request): %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		Logger.Error().Msgf("Get Url With Cookies (Status Code): " + res.Status)
		szs.reauthenticate = true
		return nil, fmt.Errorf("needs reauthentication")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		Logger.Error().Msgf("Get Url With Cookies (ReadAll): %s", err.Error())
		return nil, err
	}

	Logger.Debug().Msg("Get Url With Cookies Success")

	return body, nil

}

func (szs *SolarZeroScrapeImpl) getCurrentData() bool {
	Logger.Info().Msg("Get Current Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getCurrentData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		Logger.Error().Msgf("Get Current Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.currentData, err = jsontypes.UnmarshalCurrentData(body)
	if err != nil {
		Logger.Error().Msgf("Get Current Data (UnmarshalCurrentData): %s", err.Error())
		return false
	}
	Logger.Debug().RawJSON("CurrentData", body)
	Logger.Info().Msg("Get Current Data Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getDayData() bool {
	Logger.Info().Msg("Get Day Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getDayData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		Logger.Error().Msgf("Get Day Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.dayData, err = jsontypes.UnmarshalDayData(body)
	if err != nil {
		Logger.Error().Msgf("Get Day Data (UnmarshalDayData): %s", err.Error())
		return false
	}
	Logger.Debug().RawJSON("DayData", body)
	Logger.Info().Msg("Get Day Data Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getMonthData() bool {
	Logger.Info().Msg("Get Month Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getMonthData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		Logger.Error().Msgf("Get Month Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.monthData, err = jsontypes.UnmarshalMonthData(body)
	if err != nil {
		Logger.Error().Msgf("Get Month Data (UnmarshalMonthData): %s", err.Error())
		return false
	}
	Logger.Debug().RawJSON("MonthData", body)
	Logger.Info().Msg("Get Month Data Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getYearData() bool {
	Logger.Info().Msg("Get Year Data")
	body, err := szs.getWithCookies(fmt.Sprintf("%s/getYearData/data?id=%s&api=%s", szs.logindata.Auth.API, szs.logindata.DeviceID.ID, szs.logindata.Auth.EMSAPI))
	if err != nil {
		Logger.Error().Msgf("Get Year Data (getWithCookies): %s", err.Error())
		return false
	}
	szs.yearData, err = jsontypes.UnmarshalYearData(body)
	if err != nil {
		Logger.Error().Msgf("Get Year Data (UnmarshalYearData): %s", err.Error())
		return false
	}
	Logger.Debug().RawJSON("YearData", body)
	Logger.Info().Msg("Get Year Data Success")
	return true
}

func (szs *SolarZeroScrapeImpl) AuthenticateFully() bool {
	if !szs.cognitoAuth() {
		return false
	}

	if !szs.fetchSalesForceData() {
		return false
	}

	return szs.getCookies()
}

func (szs *SolarZeroScrapeImpl) GetData() bool {
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

func (szs *SolarZeroScrapeImpl) CurrentData() jsontypes.CurrentData {
	return szs.currentData
}

func (szs *SolarZeroScrapeImpl) DayData() jsontypes.DayData {
	return szs.dayData
}

func (szs *SolarZeroScrapeImpl) MonthData() jsontypes.MonthData {
	return szs.monthData
}

func (szs *SolarZeroScrapeImpl) YearData() jsontypes.YearData {
	return szs.yearData
}

func (szs *SolarZeroScrapeImpl) Ready() bool {
	return szs.ready
}

func (szs *SolarZeroScrapeImpl) Healthy() bool {
	diff := time.Since(szs.lastGoodWriteTimestamp)
	return diff > 0 && diff.Minutes() <= 10
}
