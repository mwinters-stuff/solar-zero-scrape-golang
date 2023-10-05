package jsontypes_test

import (
	"testing"
	"time"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeCurrentData(t *testing.T) {
	json := `
	{
		"ppv1": 0.16,
		"ppv2": 0,
		"receivedDate": "2023-10-04T17:43:17",
		"soc": 90,
		"load": 0.94,
		"deviceStatus": 1,
		"temperature": 27.6,
		"import": 0.02,
		"export": 0,
		"batteryCurrent": 14.5,
		"batteryVoltage": 52.8,
		"charge": 0,
		"discharge": 0.76,
		"powerFlow": 7,
		"gridPowerMode": 1,
		"gridPowerOutage": 0,
		"interconnectionState": null,
		"totalCapacity": 0
	}`

	currentdata, err := jsontypes.UnmarshalCurrentData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, 0.16, currentdata.Ppv1)
	assert.Equal(t, 0.0, currentdata.Ppv2)
	assert.Equal(t, "2023-10-04T17:43:17", currentdata.ReceivedDate)
	assert.Equal(t, int64(90), currentdata.Soc)
	assert.Equal(t, 0.94, currentdata.Load)
	assert.Equal(t, int64(1), currentdata.DeviceStatus)
	assert.Equal(t, 27.6, currentdata.Temperature)
	assert.Equal(t, 0.02, currentdata.Import)
	assert.Equal(t, 0.0, currentdata.Export)
	assert.Equal(t, 14.5, currentdata.BatteryCurrent)
	assert.Equal(t, 52.8, currentdata.BatteryVoltage)
	assert.Equal(t, 0.0, currentdata.Charge)
	assert.Equal(t, 0.76, currentdata.Discharge)
	assert.Equal(t, int64(7), currentdata.PowerFlow)
	assert.Equal(t, int64(0), currentdata.GridPowerOutage)
	assert.Equal(t, int64(1), currentdata.GridPowerMode)

}

func TestGetInfluxFields(t *testing.T) {
	json := `
	{
		"ppv1": 0.16,
		"ppv2": 0,
		"receivedDate": "2023-10-04T17:43:17",
		"soc": 90,
		"load": 0.94,
		"deviceStatus": 1,
		"temperature": 27.6,
		"import": 0.02,
		"export": 0,
		"batteryCurrent": 14.5,
		"batteryVoltage": 52.8,
		"charge": 0,
		"discharge": 0.76,
		"powerFlow": 7,
		"gridPowerMode": 1,
		"gridPowerOutage": 0,
		"interconnectionState": null,
		"totalCapacity": 0
	}`

	currentdata, err := jsontypes.UnmarshalCurrentData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	influxData := currentdata.GetInfluxFields()

	assert.Equal(t, int64(1), influxData["DeviceStatus"])
	assert.Equal(t, int64(7), influxData["DPowerFlow"])
	assert.Equal(t, float64(0), influxData["Export"])
	assert.Equal(t, float64(20), influxData["Import"])
	assert.Equal(t, float64(940), influxData["Load"])
	assert.Equal(t, float64(160), influxData["Solar"])
	assert.Equal(t, int64(90), influxData["Soc"])
	assert.Equal(t, float64(0), influxData["Charge"])
	assert.Equal(t, float64(760), influxData["Discharge"])
	assert.Equal(t, int64(0), influxData["GridPowerOutage"])
	assert.Equal(t, 27.6, influxData["Temperature"])

}

func parseLocalTimestamp(timestamp string) (time.Time, error) {
	// Load the local timezone
	local, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}

	// Define the layout of the timestamp string
	layout := "2006-01-02T15:04:05"

	// Parse the timestamp string into a time.Time instance in the local timezone
	t, err := time.ParseInLocation(layout, timestamp, local)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func TestTimestampConversion(t *testing.T) {
	ReceivedDate := "2023-10-05T07:51:24"
	ReceivedDateZoned := "2023-10-05T07:51:24+13:00"

	stamp, _ := parseLocalTimestamp(ReceivedDate)
	// lx := stamp.Local().Format(time.RFC3339)
	stampZoned, _ := time.Parse(time.RFC3339, ReceivedDateZoned)

	// localstr := time.Now().Format(time.RFC3339)
	// assert.Equal(t, localstr, lx)
	assert.Equal(t, stamp, stampZoned)

}
