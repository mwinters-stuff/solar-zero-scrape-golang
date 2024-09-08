package jsontypes_test

import (
	"os"
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeConfigurationData(t *testing.T) {
	json := `
	{
		"SolarZero": {
				"Username": "your@email",
				"Password": "password"
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
				"Password": "password"
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
