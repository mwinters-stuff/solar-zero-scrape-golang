package yeardata

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type YearDataTestSuite struct {
	suite.Suite
	json string
}

func (suite *YearDataTestSuite) SetupTest() {
	suite.json = `
	[
		{
			"Month": "Jan",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Month": "Feb",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Month": "Mar",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Month": "Apr",
			"Solar use": 44.489999999999995,
			"Grid": 396.15,
			"Export": -25.11,
			"Battery grid": 0.6,
			"Home load": 440.73999999999995
		},
		{
			"Month": "May",
			"Solar use": 0,
			"Grid": 702.1,
			"Export": -4.68,
			"Battery grid": 65.2,
			"Home load": 698.32
		},
		{
			"Month": "Jun",
			"Solar use": 115.74,
			"Grid": 914.49,
			"Export": -2.36,
			"Battery grid": 279.4,
			"Home load": 1014.5299999999999
		},
		{
			"Month": "Jul",
			"Solar use": 113.51,
			"Grid": 1251.37,
			"Export": -1.69,
			"Battery grid": 309.3,
			"Home load": 1351.08
		},
		{
			"Month": "Aug",
			"Solar use": 198.75,
			"Grid": 925.29,
			"Export": -9.55,
			"Battery grid": 274.6,
			"Home load": 1110.54
		},
		{
			"Month": "Sep",
			"Solar use": 62.300000000000004,
			"Grid": 201.91,
			"Export": -6.9,
			"Battery grid": 68.3,
			"Home load": 261.41
		},
		{
			"Month": "Oct",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Month": "Nov",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		},
		{
			"Month": "Dec",
			"Solar use": null,
			"Grid": 0,
			"Export": 0,
			"Battery grid": null,
			"Home load": null
		}
	]`

}

func (suite *YearDataTestSuite) TestUnmarshal() {

	yeardata, err := UnmarshalYearData([]byte(suite.json))
	assert.Nil(suite.T(), err, "Err is not nil")

	assert.Len(suite.T(), yeardata, 12, "Year data must have 12 items")

	month0 := yeardata[0]
	assert.Equal(suite.T(), "Jan", month0.Month)
	assert.Nil(suite.T(), month0.SolarUse)
	assert.Equal(suite.T(), float64(0), *month0.Export)
	assert.Equal(suite.T(), float64(0), *month0.Grid)
	assert.Nil(suite.T(), month0.BatteryGrid)
	assert.Nil(suite.T(), month0.HomeLoad)

	month7 := yeardata[6]
	assert.Equal(suite.T(), "Jul", month7.Month)
	assert.Equal(suite.T(), float64(113.51), *month7.SolarUse)
	assert.Equal(suite.T(), float64(-1.69), *month7.Export)
	assert.Equal(suite.T(), float64(1251.37), *month7.Grid)
	assert.Equal(suite.T(), float64(309.3), *month7.BatteryGrid)
	assert.Equal(suite.T(), float64(1351.08), *month7.HomeLoad)

	influx0 := month0.GetInfluxFields()
	assert.Nil(suite.T(), influx0)

	influx7 := month7.GetInfluxFields()
	assert.NotNil(suite.T(), influx7)

	assert.Equal(suite.T(), "Jul", (*influx7)["Month"])
	assert.Equal(suite.T(), float64(113.51), (*influx7)["SolarUse"])
	assert.Equal(suite.T(), float64(-1.69), (*influx7)["Export"])
	assert.Equal(suite.T(), float64(1251.37), (*influx7)["Grid"])
	assert.Equal(suite.T(), float64(309.3), (*influx7)["BatteryGrid"])
	assert.Equal(suite.T(), float64(1351.08), (*influx7)["HomeLoad"])
}

func (suite *YearDataTestSuite) TestMonthNumber() {

	yeardata, err := UnmarshalYearData([]byte(suite.json))
	assert.Nil(suite.T(), err, "Err is not nil")

	for i, yd := range yeardata {
		assert.Equal(suite.T(), time.Month(i+1), yd.GetMonthNum())
	}
}

func TestYearTestDataSuite(t *testing.T) {
	suite.Run(t, new(YearDataTestSuite))
}
