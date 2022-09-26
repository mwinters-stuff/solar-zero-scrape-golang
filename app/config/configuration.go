// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    configuration, err := UnmarshalConfiguration(bytes)
//    bytes, err = configuration.Marshal()

package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

func UnmarshalConfiguration(data []byte) (Configuration, error) {
	var r Configuration
	err := json.Unmarshal(data, &r)
	return r, err
}

func LoadConfiguration(filename string) (Configuration, error) {
	var c Configuration
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Error().Err(err)
		return c, err
	}

	c, err = UnmarshalConfiguration(content)
	return c, err
}

type Configuration struct {
	SolarZero SolarZero `json:"SolarZero"`
	InfluxDB  InfluxDB  `json:"InfluxDB"`
	Mqtt      Mqtt      `json:"MQTT"`
}

type InfluxDB struct {
	Token   string `json:"Token"`
	HostURL string `json:"HostUrl"`
	Org     string `json:"Org"`
	Bucket  string `json:"Bucket"`
}

type Mqtt struct {
	URL       string `json:"URL"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	BaseTopic string `json:"BaseTopic"`
}

type SolarZero struct {
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	UserPoolID string `json:"UserPoolId"`
	ClientID   string `json:"ClientId"`
	API        API    `json:"API"`
}

type API struct {
	Region              string `json:"Region"`
	APIGatewayURL       string `json:"ApiGatewayURL"`
	APIKey              string `json:"ApiKey"`
	SolarZeroAPIAddress string `json:"SolarZeroApiAddress"`
}
