package rangedata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeRangeData(t *testing.T) {
	json := `
	[
		{
			"Date": "1-Jul",
			"Solar use": 6.029999999999999,
			"Grid": 42.71,
			"Export": -0.07,
			"Load": 48.34,
			"Solar": 6.1,
			"Discharge": 9.4,
			"Charge": 9.8
		},
		{
			"Date": "2-Jul",
			"Solar use": 5.63,
			"Grid": 40.06,
			"Export": -0.07,
			"Load": 45.29,
			"Solar": 5.7,
			"Discharge": 9.4,
			"Charge": 9.8
		},
		{
			"Date": "3-Jul",
			"Solar use": 6.05,
			"Grid": 39.58,
			"Export": -0.05,
			"Load": 45.23,
			"Solar": 6.1,
			"Discharge": 9.4,
			"Charge": 9.8
		},
		{
			"Date": "4-Jul",
			"Solar use": 5.78,
			"Grid": 34.19,
			"Export": -0.12,
			"Load": 39.67,
			"Solar": 5.9,
			"Discharge": 9.5,
			"Charge": 9.8
		},
		{
			"Date": "5-Jul",
			"Solar use": 1.89,
			"Grid": 39.69,
			"Export": -0.01,
			"Load": 41.08,
			"Solar": 1.9,
			"Discharge": 9.5,
			"Charge": 10
		},
		{
			"Date": "6-Jul",
			"Solar use": 6.029999999999999,
			"Grid": 37.2,
			"Export": -0.07,
			"Load": 42.83,
			"Solar": 6.1,
			"Discharge": 9.8,
			"Charge": 10.2
		},
		{
			"Date": "7-Jul",
			"Solar use": 2.5700000000000003,
			"Grid": 48.47,
			"Export": -0.03,
			"Load": 50.54,
			"Solar": 2.6,
			"Discharge": 9.7,
			"Charge": 10.2
		},
		{
			"Date": "8-Jul",
			"Solar use": 1.29,
			"Grid": 42.72,
			"Export": -0.01,
			"Load": 43.61,
			"Solar": 1.3,
			"Discharge": 9.6,
			"Charge": 10
		}
	]
	`

	data, err := UnmarshalRangeExportData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Len(t, data, 8, "Range data must have 8 items")

	day := data[0]
	assert.Equal(t, "1-Jul", day.Date)
	assert.Equal(t, 6.029999999999999, day.SolarUse)
	assert.Equal(t, 42.71, day.Grid)
	assert.Equal(t, -0.07, day.Export)
	assert.Equal(t, 48.34, day.Load)
	assert.Equal(t, 6.1, day.Solar)
	assert.Equal(t, 9.4, day.Discharge)
	assert.Equal(t, 9.8, day.Charge)

	day = data[1]
	assert.Equal(t, "2-Jul", day.Date)
	assert.Equal(t, 5.63, day.SolarUse)
	assert.Equal(t, 40.06, day.Grid)
	assert.Equal(t, 0.07, day.Export)
	assert.Equal(t, 45.29, day.Load)
	assert.Equal(t, 5.7, day.Solar)
	assert.Equal(t, 9.4, day.Discharge)
	assert.Equal(t, 9.8, day.Charge)
}
