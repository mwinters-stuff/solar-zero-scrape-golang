package jsontypes_test

import (
	"os"
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeConfigurationData(t *testing.T) {
	json := `
	{
		"SolarZero": {
				"Username": "your@email",
				"Password": "password",
				"UserPoolId": "us-west-2_NoMpv1v1A",
				"ClientId": "6mgtqq7vvf7eo3r3qrsg6kl1tf",
				"API": {
					"Region" :"us-west-2",
					"ApiGatewayURL": "https://d6nfzye2cb.execute-api.us-west-2.amazonaws.com",
					"ApiKey": "mA0UW2ldUUQBY3e9bZWq9lCeKQUNCZC9oKidvdbb",
					"SolarZeroApiAddress": "solarzero.pnz.technology"
				}
		},
		"InfluxDB":{
				"Token": "token",
				"HostUrl": "https://influxdb.example.com",
				"Org": "example.com",
				"Bucket": "solarzero/autogen"
		},
		"MQTT": {
			"URL":"mqtt://example.com:1883",
			"Username": "solarzero",
			"Password": "zerosolar",
			"BaseTopic":"solar-zero"
		}
	}
	`

	data, err := jsontypes.UnmarshalConfiguration([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "your@email", data.SolarZero.Username)
	assert.Equal(t, "password", data.SolarZero.Password)
	assert.Equal(t, "us-west-2_NoMpv1v1A", data.SolarZero.UserPoolID)
	assert.Equal(t, "6mgtqq7vvf7eo3r3qrsg6kl1tf", data.SolarZero.ClientID)
	assert.Equal(t, "us-west-2", data.SolarZero.API.Region)
	assert.Equal(t, "https://d6nfzye2cb.execute-api.us-west-2.amazonaws.com", data.SolarZero.API.APIGatewayURL)
	assert.Equal(t, "mA0UW2ldUUQBY3e9bZWq9lCeKQUNCZC9oKidvdbb", data.SolarZero.API.APIKey)
	assert.Equal(t, "solarzero.pnz.technology", data.SolarZero.API.SolarZeroAPIAddress)

	assert.Equal(t, "token", data.InfluxDB.Token)
	assert.Equal(t, "https://influxdb.example.com", data.InfluxDB.HostURL)
	assert.Equal(t, "example.com", data.InfluxDB.Org)
	assert.Equal(t, "solarzero/autogen", data.InfluxDB.Bucket)

	assert.Equal(t, "mqtt://example.com:1883", data.Mqtt.URL)
	assert.Equal(t, "solarzero", data.Mqtt.Username)
	assert.Equal(t, "zerosolar", data.Mqtt.Password)
	assert.Equal(t, "solar-zero", data.Mqtt.BaseTopic)

}

func TestLoadConfiguration(t *testing.T) {
	//write temporary file
	json := `
	{
		"SolarZero": {
				"Username": "your@email",
				"Password": "password",
				"UserPoolId": "us-west-2_NoMpv1v1A",
				"ClientId": "6mgtqq7vvf7eo3r3qrsg6kl1tf",
				"API": {
					"Region" :"us-west-2",
					"ApiGatewayURL": "https://d6nfzye2cb.execute-api.us-west-2.amazonaws.com",
					"ApiKey": "mA0UW2ldUUQBY3e9bZWq9lCeKQUNCZC9oKidvdbb",
					"SolarZeroApiAddress": "solarzero.pnz.technology"
				}
		},
		"InfluxDB":{
				"Token": "token",
				"HostUrl": "https://influxdb.example.com",
				"Org": "example.com",
				"Bucket": "solarzero/autogen"
		},
		"MQTT": {
			"URL":"mqtt://example.com:1883",
			"Username": "solarzero",
			"Password": "zerosolar",
			"BaseTopic":"solar-zero"
		}
	}
	`
	file, err := os.CreateTemp("", "cfg")
	assert.Nil(t, err, "Err is not nil")
	file.WriteString(json)

	tempFile := file.Name()

	file.Close()

	data, err := jsontypes.LoadConfiguration(tempFile)
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, "your@email", data.SolarZero.Username)
	assert.Equal(t, "password", data.SolarZero.Password)
	assert.Equal(t, "us-west-2_NoMpv1v1A", data.SolarZero.UserPoolID)
	assert.Equal(t, "6mgtqq7vvf7eo3r3qrsg6kl1tf", data.SolarZero.ClientID)
	assert.Equal(t, "us-west-2", data.SolarZero.API.Region)
	assert.Equal(t, "https://d6nfzye2cb.execute-api.us-west-2.amazonaws.com", data.SolarZero.API.APIGatewayURL)
	assert.Equal(t, "mA0UW2ldUUQBY3e9bZWq9lCeKQUNCZC9oKidvdbb", data.SolarZero.API.APIKey)
	assert.Equal(t, "solarzero.pnz.technology", data.SolarZero.API.SolarZeroAPIAddress)

	assert.Equal(t, "token", data.InfluxDB.Token)
	assert.Equal(t, "https://influxdb.example.com", data.InfluxDB.HostURL)
	assert.Equal(t, "example.com", data.InfluxDB.Org)
	assert.Equal(t, "solarzero/autogen", data.InfluxDB.Bucket)

	assert.Equal(t, "mqtt://example.com:1883", data.Mqtt.URL)
	assert.Equal(t, "solarzero", data.Mqtt.Username)
	assert.Equal(t, "zerosolar", data.Mqtt.Password)
	assert.Equal(t, "solar-zero", data.Mqtt.BaseTopic)
	os.Remove(tempFile)

}

func TestLoadConfigurationFailed(t *testing.T) {
	data, err := jsontypes.LoadConfiguration("badfilename")
	assert.NotNil(t, err, "Err is nil")
	assert.NotNil(t, data, "Data is nil")

}
