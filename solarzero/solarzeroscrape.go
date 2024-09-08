package solarzero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

type SolarZeroScrape interface {
	Start()

	Ready() bool
	Healthy() bool

	Customer() jsontypes.CustomerData
	Data() jsontypes.DataResponseData
	Daily() jsontypes.DailyResponseData
}

type SolarZeroScrapeImpl struct {
	config jsontypes.Configuration

	influxdb      InfluxDBWriter
	mqtt          MQTTClient
	correlationID string

	authResponse jsontypes.AuthResponse
	customerData jsontypes.CustomerData
	data         jsontypes.DataResponseData
	daily        jsontypes.DailyResponseData

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
		config:                 config,
		influxdb:               influxdb,
		mqtt:                   mqtt,
		lastGoodWriteTimestamp: time.Now(),
		ready:                  false,
	}

	return scrape
}

func (szs *SolarZeroScrapeImpl) Start() {
	szs.ready = true

	s := gocron.NewScheduler(time.Local)

	if err := szs.login(szs.config.SolarZero.Username, szs.config.SolarZero.Password); err != nil {
		Logger.Error().Msgf("AuthenicateFully Failed: %s", err)
		os.Exit(-1)
	}

	err := szs.getCustomer()
	if err != nil {
		Logger.Error().Msg(err.Error())
		os.Exit(-1)
	}

	s.Every(1).Minutes().Do(func() {
		Logger.Info().Msgf("Get info @ %s", time.Now().String())
		err := szs.getData()
		if err != nil {
			Logger.Error().Msg(err.Error())
		}
		if err == nil && szs.mqtt != nil {
			szs.mqtt.WriteData(szs)
		}
	})

	s.Every(1).Days().Do(func() {
		Logger.Info().Msgf("Get daily @ %s", time.Now().String())
		err := szs.getDaily()
		if err != nil {
			Logger.Error().Msg(err.Error())
		}
		err = szs.getCustomer()
		if err != nil {
			Logger.Error().Msg(err.Error())
		}

	})

	s.Every(1).Hours().Do(func() {
		if szs.mqtt != nil {
			Logger.Info().Msgf("Publish HomeAssistant @ %s", time.Now().String())
			szs.mqtt.PublishHomeAssistantDiscovery()
		}
	})

	s.StartBlocking()
}

func (szs *SolarZeroScrapeImpl) login(username, password string) error {

	szs.correlationID = uuid.New().String()

	url := httpURL + httpURLAuth
	body := map[string]string{
		"username": username,
		"password": password,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Correlation-Id", szs.correlationID)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	szs.authResponse, err = jsontypes.UnmarshalAuthResponse(bodyBytes)
	if err != nil {
		return err
	}

	Logger.Info().Msg("Logged in successfully!")
	return nil
}

func (szs *SolarZeroScrapeImpl) refreshAccessToken() error {
	url := httpURL + httpURLRefresh
	body := map[string]string{
		"refreshToken": szs.authResponse.Tokens.RefreshToken,
		"idToken":      szs.authResponse.Tokens.IDToken,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Correlation-Id", szs.correlationID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token refresh failed: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	szs.authResponse, err = jsontypes.UnmarshalAuthResponse(bodyBytes)
	if err != nil {
		return err
	}

	Logger.Info().Msg("Token refreshed successfully!")
	return nil

}

func (szs *SolarZeroScrapeImpl) makeAuthenticatedGETRequest(requests string) (string, error) {
	if time.Now().After(szs.authResponse.Tokens.ExpiresAt) {
		// Access token has expired, refresh it
		err := szs.refreshAccessToken()
		if err != nil {
			return "", err
		}
	}

	url := httpURL + requests
	Logger.Info().Msgf("GET %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set Authorization header with the current access token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+szs.authResponse.Tokens.IDToken)
	req.Header.Set("x-correlation-id", szs.correlationID)
	req.Header.Set("x-session-id", szs.authResponse.Tokens.SessionID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println("Response:", string(bodyBytes))
	return string(bodyBytes), nil
}

func (szs *SolarZeroScrapeImpl) makeAuthenticatedPOSTRequest(requests string, body []byte) (string, error) {
	if time.Now().After(szs.authResponse.Tokens.ExpiresAt) {
		// Access token has expired, refresh it
		err := szs.refreshAccessToken()
		if err != nil {
			return "", err
		}
	}

	url := httpURL + requests
	Logger.Info().Msgf("GET %s with %s", url, string(body))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	// Set Authorization header with the current access token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+szs.authResponse.Tokens.IDToken)
	req.Header.Set("x-correlation-id", szs.correlationID)
	req.Header.Set("x-session-id", szs.authResponse.Tokens.SessionID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// fmt.Println("Response from /data:", string(bodyBytes))
	return string(bodyBytes), nil
}

func (szs *SolarZeroScrapeImpl) getCustomer() error {
	var err error
	var response = ""
	if response, err = szs.makeAuthenticatedGETRequest(httpURLCustomers); err != nil {
		Logger.Error().Msgf("Get Customer Info Failed: %s", err)
		os.Exit(-1)
	}

	if szs.customerData, err = jsontypes.UnmarshalCustomerData([]byte(response)); err != nil {
		Logger.Error().Msgf("Get Customer Info Failed: %s", err)
		os.Exit(-1)
	}
	return nil
}

func (szs *SolarZeroScrapeImpl) getData() error {

	var infoRequestData jsontypes.DataRequestData
	infoRequestData.HasTou = true
	infoRequestData.ProviderID = "nz-ecotricity"
	infoRequestData.Timezone = "Pacific/Auckland"
	infoRequestData.SiteID = szs.customerData.Account.SiteID
	bytes, _ := infoRequestData.Marshal()

	var err error
	var response = ""

	if response, err = szs.makeAuthenticatedPOSTRequest(httpURLInfo, bytes); err != nil {
		Logger.Error().Msgf("Get Data Failed: %s", err)
		return err
	}

	if szs.data, err = jsontypes.UnmarshalDataResponseData([]byte(response)); err != nil {
		Logger.Error().Msgf("Get Data Info Failed: %s", err)
		return err
	}
	return nil
}

func (szs *SolarZeroScrapeImpl) getDaily() error {

	var dailyRequestData jsontypes.DailyRequestData
	dailyRequestData.HasTou = true
	dailyRequestData.Timezone = "Pacific/Auckland"
	dailyRequestData.SiteID = szs.customerData.Account.SiteID
	bytes, _ := dailyRequestData.Marshal()

	var err error
	var response = ""

	if response, err = szs.makeAuthenticatedPOSTRequest(httpURLDaily, bytes); err != nil {
		Logger.Error().Msgf("Get Daily Failed: %s", err)
		return err
	}

	if szs.daily, err = jsontypes.UnmarshalDailyResponseData([]byte(response)); err != nil {
		Logger.Error().Msgf("Get Daily Info Failed: %s", err)
		return err
	}
	return nil
}

func (szs *SolarZeroScrapeImpl) Daily() jsontypes.DailyResponseData {
	return szs.daily
}

func (szs *SolarZeroScrapeImpl) Data() jsontypes.DataResponseData {
	return szs.data
}

func (szs *SolarZeroScrapeImpl) Customer() jsontypes.CustomerData {
	return szs.customerData
}

func (szs *SolarZeroScrapeImpl) Ready() bool {
	return szs.ready
}

func (szs *SolarZeroScrapeImpl) Healthy() bool {
	diff := time.Since(szs.lastGoodWriteTimestamp)
	return diff > 0 && diff.Minutes() <= 10
}
