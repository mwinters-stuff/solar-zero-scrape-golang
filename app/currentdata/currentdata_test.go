package currentdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCurrentData(t *testing.T) {
	json := `
	{
		"deviceStatus": 1,
		"dPowerFlow": 5,
		"export": 0,
		"import": 1026,
		"load": 1026,
		"solar": 0,
		"soc": 19,
		"charge": 0,
		"gridPowerOutage": 0,
		"temperature": 23.1
	}`

	currentdata, err := UnmarshalCurrentData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, int64(1), currentdata.DeviceStatus)
	assert.Equal(t, int64(5), currentdata.DPowerFlow)
	assert.Equal(t, int64(0), currentdata.Export)
	assert.Equal(t, int64(1026), currentdata.Import)
	assert.Equal(t, int64(1026), currentdata.Load)
	assert.Equal(t, int64(0), currentdata.Solar)
	assert.Equal(t, int64(19), currentdata.Soc)
	assert.Equal(t, int64(0), currentdata.Charge)
	assert.Equal(t, int64(0), currentdata.GridPowerOutage)
	assert.Equal(t, 23.1, currentdata.Temperature)

}

func TestGetInfluxFields(t *testing.T) {
	currentdata := CurrentData{
		DeviceStatus:    1,
		DPowerFlow:      5,
		Export:          0,
		Import:          1026,
		Load:            1026,
		Solar:           0,
		Soc:             19,
		Charge:          0,
		GridPowerOutage: 0,
		Temperature:     23.1,
	}

	influxData := currentdata.GetInfluxFields()

	assert.Equal(t, int64(1), influxData["DeviceStatus"])
	assert.Equal(t, int64(5), influxData["DPowerFlow"])
	assert.Equal(t, int64(0), influxData["Export"])
	assert.Equal(t, int64(1026), influxData["Import"])
	assert.Equal(t, int64(1026), influxData["Load"])
	assert.Equal(t, int64(0), influxData["Solar"])
	assert.Equal(t, int64(19), influxData["Soc"])
	assert.Equal(t, int64(0), influxData["Charge"])
	assert.Equal(t, int64(0), influxData["GridPowerOutage"])
	assert.Equal(t, 23.1, influxData["Temperature"])

}
