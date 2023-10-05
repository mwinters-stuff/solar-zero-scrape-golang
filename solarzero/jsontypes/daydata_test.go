package jsontypes_test

import (
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeDayData(t *testing.T) {
	json := `[
		{
			"receivedDate": "2023-10-04T00:00:38",
			"solarUse": 0,
			"grid": 0.68,
			"export": 0,
			"stateOfCharge": 20
		},
		{
			"receivedDate": "2023-10-04T01:04:01",
			"solarUse": 0,
			"grid": 1.8,
			"export": 0,
			"stateOfCharge": 30
		},
		{
			"receivedDate": "2023-10-04T02:02:33",
			"solarUse": 0,
			"grid": 1.8,
			"export": 0,
			"stateOfCharge": 50
		},
		{
			"receivedDate": "2023-10-04T03:01:04",
			"solarUse": 0,
			"grid": 1.76,
			"export": 0,
			"stateOfCharge": 71
		},
		{
			"receivedDate": "2023-10-04T04:04:28",
			"solarUse": 0,
			"grid": 2.62,
			"export": 0,
			"stateOfCharge": 90
		},
		{
			"receivedDate": "2023-10-04T05:02:58",
			"solarUse": 0,
			"grid": 0.94,
			"export": 0,
			"stateOfCharge": 97
		},
		{
			"receivedDate": "2023-10-04T06:01:30",
			"solarUse": 0.01,
			"grid": 1.52,
			"export": 0,
			"stateOfCharge": 97
		},
		{
			"receivedDate": "2023-10-04T07:00:01",
			"solarUse": 0.04,
			"grid": 0.25,
			"export": -0.01,
			"stateOfCharge": 77
		},
		{
			"receivedDate": "2023-10-04T08:03:25",
			"solarUse": 0.1,
			"grid": 0.01,
			"export": -0.06,
			"stateOfCharge": 39
		},
		{
			"receivedDate": "2023-10-04T09:01:54",
			"solarUse": 0.22,
			"grid": 0.55,
			"export": -0.01,
			"stateOfCharge": 21
		},
		{
			"receivedDate": "2023-10-04T10:00:23",
			"solarUse": 0.8,
			"grid": 0.12,
			"export": -0.02,
			"stateOfCharge": 21
		},
		{
			"receivedDate": "2023-10-04T11:03:45",
			"solarUse": 1.47,
			"grid": 0.01,
			"export": -0.02,
			"stateOfCharge": 24
		},
		{
			"receivedDate": "2023-10-04T12:02:14",
			"solarUse": 1.96,
			"grid": 0.01,
			"export": -0.03,
			"stateOfCharge": 28
		},
		{
			"receivedDate": "2023-10-04T13:00:43",
			"solarUse": 2.32,
			"grid": 0,
			"export": -0.01,
			"stateOfCharge": 50
		},
		{
			"receivedDate": "2023-10-04T14:04:03",
			"solarUse": 1.7,
			"grid": 0.06,
			"export": -0.55,
			"stateOfCharge": 71
		},
		{
			"receivedDate": "2023-10-04T15:02:30",
			"solarUse": 1.82,
			"grid": 0.11,
			"export": -0.28,
			"stateOfCharge": 90
		},
		{
			"receivedDate": "2023-10-04T16:00:57",
			"solarUse": 0.87,
			"grid": 0.32,
			"export": -0.29,
			"stateOfCharge": 97
		},
		{
			"receivedDate": "2023-10-04T17:04:18",
			"solarUse": 0.22,
			"grid": 0.02,
			"export": 0,
			"stateOfCharge": 94
		}
	]`

	daydata, err := jsontypes.UnmarshalDayData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Len(t, daydata, 18, "Day data must have 18 items")

	hour0 := daydata[15]

	assert.Equal(t, "2023-10-04T15:02:30", hour0.ReceivedDate)
	assert.Equal(t, float64(-0.28), hour0.Export)
	assert.Equal(t, float64(0.11), hour0.Grid)
	assert.Equal(t, float64(1.820), hour0.SolarUse)
	assert.Equal(t, int64(90), hour0.StateOfCharge)

	influx0 := hour0.GetInfluxFields()
	assert.NotNil(t, influx0)
	assert.Equal(t, "2023-10-04T15:02:30", (*influx0)["Hour"])
	assert.Equal(t, float64(-0.28), (*influx0)["Export"])
	assert.Equal(t, float64(0.11), (*influx0)["Grid"])
	assert.Equal(t, float64(1.820), (*influx0)["SolarUse"])
	assert.Equal(t, int64(90), (*influx0)["SoC"])

}
