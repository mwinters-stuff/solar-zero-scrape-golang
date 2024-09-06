package solarzero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

type SolarZeroScrape interface {
	Start()
	AuthenticateFully() bool
	GetData() bool
	CurrentData() jsontypes.CurrentData
	DayData() jsontypes.DayData
	MonthData() jsontypes.DayData
	YearData() jsontypes.DayData

	SolarVsGrid() jsontypes.SolarVsGrid
	SolarUse() jsontypes.SolarUse
	ElectricityUse() jsontypes.ElectricityUse

	Ready() bool
	Healthy() bool
}

type SolarZeroScrapeImpl struct {
	config jsontypes.Configuration
	// accessToken  string
	// refreshToken string
	// idToken      string
	apiRequestURL string
	apiToken      string
	customerID    string
	customerUUID  string

	// userAttributes map[string]string
	salesForceData jsontypes.SalesForceData
	cookies        []*http.Cookie
	reauthenticate bool

	influxdb InfluxDBWriter
	mqtt     MQTTClient

	currentData jsontypes.CurrentData
	dayData     jsontypes.DayData
	monthData   jsontypes.DayData
	yearData    jsontypes.DayData

	solarVsGrid    jsontypes.SolarVsGrid
	solarUse       jsontypes.SolarUse
	electricityUse jsontypes.ElectricityUse

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

	var influxdb InfluxDBWriter
	if config.InfluxDB.HostURL != "" {
		influxdb = NewInfluxDBWriter(&config)
		var err = influxdb.Connect(influxdb2.NewClient(config.InfluxDB.HostURL, config.InfluxDB.Token))
		if err != nil {
			Logger.Panic().Msgf("InfluxDB Connect %s", err.Error())
		}
	}

	var mqtt MQTTClient
	if config.Mqtt.URL != "" {
		mqtt = NewMQTTClient(&config)
		var err = mqtt.Connect()
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
		s.Every(1).Minutes().Do(func() {
			Logger.Info().Msgf("Get Data at %s", time.Now())
			success := szs.GetData()
			if success {
				if szs.influxdb != nil {
					szs.influxdb.WriteData(szs)
				}
				if szs.mqtt != nil {
					szs.mqtt.WriteData(szs)
				}
				szs.lastGoodWriteTimestamp = time.Now()
			} else {
				Logger.Error().Msg("GetData Failed, Reauthenticating")
				go s.Stop()
			}
		})
		s.Every(1).Day().At("01:00").Do(func() {
			Logger.Info().Msgf("Get Daily Data at %s", time.Now())
			success := szs.GetDailyData()
			if success {
				if szs.influxdb != nil {
					szs.influxdb.WriteDailyData(szs)
				}
				szs.lastGoodWriteTimestamp = time.Now()
			} else {
				Logger.Error().Msg("Daily Failed, Reauthenticating")
				go s.Stop()
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

	var jsonStr, _ = json.Marshal(map[string]string{"contactId": szs.awsInterface.UserAttributes()["custom:contactId"]})

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

	//"{\"message\": \"Endpoint request timed out\"}"
	// Logger.Debug().Msgf("SalesForceData: %s", body)
	if string(body) == "{\"message\": \"Endpoint request timed out\"}" {
		Logger.Error().Msgf("SalesForceData: Endpoint request timed out")
		return false
	}

	szs.salesForceData, err = jsontypes.UnmarshalSalesForceData(body)
	if err != nil {
		Logger.Error().Msgf("Fetch SalesForce Data (Unmarshal): %s", err.Error())
		return false
	}
	Logger.Info().Msg("Fetch SalesForce Data Success")

	return true
}

func (szs *SolarZeroScrapeImpl) httpGet(url string, addCookies bool, addToken bool) ([]byte, *http.Response, error) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", url, nil)

	if addCookies {
		for _, cookie := range szs.cookies {
			req.AddCookie(cookie)
		}

	}

	if addToken {
		req.Header.Add("X-Token", szs.apiToken)
	}

	res, err := client.Do(req)
	if err != nil {
		Logger.Error().Msgf("httpGet Do (%s): %s", url, err.Error())
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		Logger.Error().Msgf("httpGet ReadAll (%s): %s", url, err.Error())
		return nil, nil, err
	}

	return body, res, nil
}

func (szs *SolarZeroScrapeImpl) getAPIAuthentication() bool {
	szs.reauthenticate = false

	Logger.Info().Msg("getAPIAuthentication - (login)")

	url := fmt.Sprintf("https://%s/login/%s",
		szs.config.SolarZero.API.SolarZeroAPIAddress, szs.salesForceData.Token)

	_, res, err := szs.httpGet(url, false, false)
	if err != nil {
		return false
	}

	szs.cookies = res.Cookies()

	if szs.cookies == nil {
		Logger.Error().Msgf("getAPIAuthentication (No Cookies)")
		return false
	}

	Logger.Info().Msg("getAPIAuthentication - (cookie)")

	url = fmt.Sprintf("https://%s/cookie", szs.config.SolarZero.API.SolarZeroAPIAddress)
	bodyBytes, _, err := szs.httpGet(url, true, false)
	if err != nil {
		return false
	}
	cookieResult, err := jsontypes.UnmarshalCookieResult(bodyBytes)
	if err != nil {
		Logger.Error().Msgf("getAPIAuthentication - (unmarshal cookieResult): %s", err.Error())
		return false
	}
	szs.customerID = cookieResult.CustomerID

	Logger.Info().Msg("getAPIAuthentication - (authentication)")

	url = fmt.Sprintf("https://%s/api/authentication", szs.config.SolarZero.API.SolarZeroAPIAddress)

	bodyBytes, _, err = szs.httpGet(url, true, false)
	if err != nil {
		return false
	}
	authResult, err := jsontypes.UnmarshalAuthResult(bodyBytes)
	if err != nil {
		Logger.Error().Msgf("getAPIAuthentication - (unmarshal authResult): %s", err.Error())
		return false
	}

	szs.apiRequestURL = authResult.URL
	szs.apiToken = authResult.TokenString

	Logger.Info().Msg("getAPIAuthentication - (findCustomerId)")

	url = fmt.Sprintf("%s/EnergyDataDevice/find?customerId=%s", szs.apiRequestURL, szs.customerID)

	bodyBytes, _, err = szs.httpGet(url, false, true)
	if err != nil {
		return false
	}
	findCustomerResult, err := jsontypes.UnmarshalFindCustomerResult(bodyBytes)
	if err != nil {
		Logger.Error().Msgf("getAPIAuthentication - (unmarshal findCustomerResult): %s", err.Error())
		return false
	}

	szs.customerUUID = findCustomerResult.ID

	return true

}

func (szs *SolarZeroScrapeImpl) getSolarDataWithToken(query string) ([]byte, error) {

	url := fmt.Sprintf("%s/SolarData/%s/%s", szs.apiRequestURL, szs.customerUUID, query)

	bodyBytes, _, err := szs.httpGet(url, false, true)

	return bodyBytes, err

}

func (szs *SolarZeroScrapeImpl) getSolarVsGridDataWithToken(query string) ([]byte, error) {

	url := fmt.Sprintf("%s/SolarVsGrid/%s/%s", szs.apiRequestURL, szs.customerUUID, query)
	Logger.Debug().Msg(url)
	bodyBytes, _, err := szs.httpGet(url, false, true)

	return bodyBytes, err

}

func (szs *SolarZeroScrapeImpl) getSolarUseDataWithToken(query string) ([]byte, error) {

	url := fmt.Sprintf("%s/SolarUse/%s/%s", szs.apiRequestURL, szs.customerUUID, query)
	Logger.Debug().Msg(url)
	bodyBytes, _, err := szs.httpGet(url, false, true)

	return bodyBytes, err

}

func (szs *SolarZeroScrapeImpl) getElectricityUseDataWithToken(query string) ([]byte, error) {

	url := fmt.Sprintf("%s/ElectricityUse/%s/%s", szs.apiRequestURL, szs.customerUUID, query)
	Logger.Debug().Msg(url)
	bodyBytes, _, err := szs.httpGet(url, false, true)

	return bodyBytes, err

}

func (szs *SolarZeroScrapeImpl) getCurrentData() bool {
	Logger.Info().Msg("getCurrentData")
	body, err := szs.getSolarDataWithToken("now")
	if err != nil {
		Logger.Error().Msgf("getCurrentData: %s", err.Error())
		return false
	}
	szs.currentData, err = jsontypes.UnmarshalCurrentData(body)
	if err != nil {
		Logger.Error().Msgf("getCurrentData (UnmarshalCurrentData): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("CurrentData", body)
	Logger.Info().Msg("getCurrentData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getCurrentSolarVsGridData() bool {
	Logger.Info().Msg("getCurrentSolarVsGridData")
	body, err := szs.getSolarVsGridDataWithToken(fmt.Sprintf("day?day=%s", time.Now().Format("2006-01-02")))
	if err != nil {
		Logger.Error().Msgf("getCurrentSolarVsGridData: %s", err.Error())
		return false
	}
	szs.solarVsGrid, err = jsontypes.UnmarshalSolarVsGrid(body)
	if err != nil {
		Logger.Error().Msgf("getCurrentSolarVsGridData (UnmarshalSolarVsGrid): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("CurrentSolarVsGridData", body)
	Logger.Info().Msg("getCurrentSolarVsGridData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getCurrentSolarUseData() bool {
	Logger.Info().Msg("getCurrentSolarUseData")
	body, err := szs.getSolarUseDataWithToken(fmt.Sprintf("day?day=%s", time.Now().Format("2006-01-02")))
	if err != nil {
		Logger.Error().Msgf("getCurrentSolarUseData: %s", err.Error())
		return false
	}
	szs.solarUse, err = jsontypes.UnmarshalSolarUse(body)
	if err != nil {
		Logger.Error().Msgf("getCurrentSolarUseData (UnmarshalSolarUse): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("CurrentSolarUseData", body)
	Logger.Info().Msg("getCurrentSolarUseData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getCurrentElectricityUseData() bool {
	Logger.Info().Msg("getCurrentElectricityUseData")
	body, err := szs.getElectricityUseDataWithToken(fmt.Sprintf("day?day=%s", time.Now().Format("2006-01-02")))
	if err != nil {
		Logger.Error().Msgf("getCurrentElectricityUseData: %s", err.Error())
		return false
	}
	szs.electricityUse, err = jsontypes.UnmarshalElectricityUse(body)
	if err != nil {
		Logger.Error().Msgf("getCurrentElectricityUseData (UnmarshalElectricityUse): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("CurrentElectricityUseData", body)
	Logger.Info().Msg("getCurrentElectricityUseData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getDayData() bool {
	Logger.Info().Msg("getDayData")
	body, err := szs.getSolarDataWithToken("today")
	if err != nil {
		Logger.Error().Msgf("getDayData: %s", err.Error())
		return false
	}
	szs.dayData, err = jsontypes.UnmarshalDayData(body)
	if err != nil {
		Logger.Error().Msgf("getDayData (UnmarshalDayData): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("DayData", body)
	Logger.Info().Msg("getDayData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getMonthData() bool {
	Logger.Info().Msg("Get Month Data")
	body, err := szs.getSolarDataWithToken("month")
	if err != nil {
		Logger.Error().Msgf("getMonthData: %s", err.Error())
		return false
	}
	szs.monthData, err = jsontypes.UnmarshalDayData(body)
	if err != nil {
		Logger.Error().Msgf("getMonthData (UnmarshalMonthData): %s %s", body, err.Error())
		szs.reauthenticate = true
		return false
	}
	Logger.Debug().RawJSON("getMonthData", body)
	Logger.Info().Msg("getMonthData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) getYearData() bool {
	Logger.Info().Msg("getYearData")
	body, err := szs.getSolarDataWithToken("year")
	if err != nil {
		Logger.Error().Msgf("getYearData: %s", err.Error())
		return false
	}
	szs.yearData, err = jsontypes.UnmarshalDayData(body)
	if err != nil {
		szs.reauthenticate = true
		Logger.Error().Msgf("getYearData (UnmarshalMonthData): %s %s", body, err.Error())
		return false
	}
	Logger.Debug().RawJSON("getYearData", body)
	Logger.Info().Msg("getYearData Success")
	return true
}

func (szs *SolarZeroScrapeImpl) AuthenticateFully() bool {
	if !szs.cognitoAuth() {
		return false
	}

	if !szs.fetchSalesForceData() {
		return false
	}

	return szs.getAPIAuthentication()
}

func (szs *SolarZeroScrapeImpl) GetData() bool {
	success := szs.getCurrentData()
	if !success {
		if szs.reauthenticate {
			if szs.getAPIAuthentication() {
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

	if !szs.getCurrentSolarUseData() {
		if szs.reauthenticate {
			if szs.getAPIAuthentication() {
				if !szs.getCurrentSolarUseData() {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if !szs.getCurrentElectricityUseData() {
		if szs.reauthenticate {
			if szs.getAPIAuthentication() {
				if !szs.getCurrentElectricityUseData() {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if !szs.getCurrentSolarVsGridData() {
		if szs.reauthenticate {
			if szs.getAPIAuthentication() {
				if !szs.getCurrentSolarVsGridData() {
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
			if szs.getAPIAuthentication() {
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

	return true
}

func (szs *SolarZeroScrapeImpl) GetDailyData() bool {
	if !szs.getMonthData() {
		if szs.reauthenticate {
			if szs.getAPIAuthentication() {
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
			if szs.getAPIAuthentication() {
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

func (szs *SolarZeroScrapeImpl) MonthData() jsontypes.DayData {
	return szs.monthData
}

func (szs *SolarZeroScrapeImpl) YearData() jsontypes.DayData {
	return szs.yearData
}

func (szs *SolarZeroScrapeImpl) SolarVsGrid() jsontypes.SolarVsGrid {
	return szs.solarVsGrid
}

func (szs *SolarZeroScrapeImpl) SolarUse() jsontypes.SolarUse {
	return szs.solarUse
}

func (szs *SolarZeroScrapeImpl) ElectricityUse() jsontypes.ElectricityUse {
	return szs.electricityUse
}

func (szs *SolarZeroScrapeImpl) Ready() bool {
	return szs.ready
}

func (szs *SolarZeroScrapeImpl) Healthy() bool {
	diff := time.Since(szs.lastGoodWriteTimestamp)
	return diff > 0 && diff.Minutes() <= 10
}
