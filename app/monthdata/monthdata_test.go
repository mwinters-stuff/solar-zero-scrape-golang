package monthdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMonthData(t *testing.T) {
	json := `
	[
		{
			"Day": 1,
			"Solar use": 7.1899999999999995,
			"Grid": 16.6,
			"Export": -2.31,
			"Solar": 9.5,
			"Battery grid": 6.3,
			"Home load": 23.490000000000002
		},
		{
			"Day": 2,
			"Solar use": 2.6700000000000004,
			"Grid": 22.93,
			"Export": -0.03,
			"Solar": 2.7,
			"Battery grid": 9.2,
			"Home load": 25.2
		},
		{
			"Day": 3,
			"Solar use": 8.49,
			"Grid": 12.64,
			"Export": -0.91,
			"Solar": 9.4,
			"Battery grid": 7.6,
			"Home load": 20.83
		},
		{
			"Day": 4,
			"Solar use": 8.89,
			"Grid": 24.88,
			"Export": -2.11,
			"Solar": 11,
			"Battery grid": 6.8,
			"Home load": 33.37
		},
		{
			"Day": 5,
			"Solar use": 5.5,
			"Grid": 33.15,
			"Export": -0.1,
			"Solar": 5.6,
			"Battery grid": 10.1,
			"Home load": 38.35
		},
		{
			"Day": 6,
			"Solar use": 9.49,
			"Grid": 37.63,
			"Export": -0.61,
			"Solar": 10.1,
			"Battery grid": 9.5,
			"Home load": 46.82000000000001
		},
		{
			"Day": 7,
			"Solar use": 9.26,
			"Grid": 36.17,
			"Export": -0.14,
			"Solar": 9.4,
			"Battery grid": 9.6,
			"Home load": 44.83
		},
		{
			"Day": 8,
			"Solar use": 10.81,
			"Grid": 16.8,
			"Export": -0.69,
			"Solar": 11.5,
			"Battery grid": 9.2,
			"Home load": 27.41
		},
		{
			"Day": 9,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 10,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 11,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 12,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 13,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 14,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 15,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 16,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 17,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 18,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 19,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 20,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 21,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 22,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 23,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 24,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 25,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 26,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 27,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 28,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 29,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Day": 30,
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Solar": null,
			"Battery grid": null,
			"Home load": null
		}
	]
	`

	monthdata, err := UnmarshalMonthData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Len(t, monthdata, 30, "Month data must have 30 items")

	day0 := monthdata[0]
	assert.Equal(t, int64(1), day0.Day)
	assert.Equal(t, float64(7.1899999999999995), *day0.SolarUse)
	assert.Equal(t, float64(16.6), *day0.Grid)
	assert.Equal(t, float64(-2.31), *day0.Export)
	assert.Equal(t, float64(9.5), *day0.Solar)
	assert.Equal(t, float64(6.3), *day0.BatteryGrid)
	assert.Equal(t, float64(23.490000000000002), *day0.HomeLoad)

	day15 := monthdata[14]
	assert.Equal(t, int64(15), day15.Day)
	assert.Nil(t, day15.SolarUse)
	assert.Equal(t, float64(0), *day15.Grid)
	assert.Equal(t, float64(0), *day15.Export)
	assert.Nil(t, day15.Solar)
	assert.Nil(t, day15.BatteryGrid)
	assert.Nil(t, day15.HomeLoad)

	influx0 := day0.GetInfluxFields()
	assert.NotNil(t, influx0)
	assert.Equal(t, int64(1), (*influx0)["Day"])
	assert.Equal(t, float64(7.1899999999999995), (*influx0)["SolarUse"])
	assert.Equal(t, float64(16.6), (*influx0)["Grid"])
	assert.Equal(t, float64(-2.31), (*influx0)["Export"])
	assert.Equal(t, float64(9.5), (*influx0)["Solar"])
	assert.Equal(t, float64(6.3), (*influx0)["BatteryGrid"])
	assert.Equal(t, float64(23.490000000000002), (*influx0)["HomeLoad"])

	influx15 := day15.GetInfluxFields()
	assert.Nil(t, influx15)
}
