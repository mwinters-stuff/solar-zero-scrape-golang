package app

import (
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/daydata"
	"github.com/stretchr/testify/assert"
)

func TestDecodeDayData(t *testing.T) {
	json := `[{"Hour":"12 am","Export":0,"Grid":1.326,"Solar use":0,"SoC":20,"Charge":0,"Discharge":0,"Solar":0,"Battery grid":0,"Home load":1.326},{"Hour":"1 am","Export":0,"Grid":3.4410000000000003,"Solar use":0,"SoC":35.916666666666664,"Charge":1.995,"Discharge":0,"Solar":0,"Battery grid":1.995,"Home load":1.445},{"Hour":"2 am","Export":0,"Grid":3.369,"Solar use":0,"SoC":71.46153846153847,"Charge":1.994,"Discharge":0,"Solar":0,"Battery grid":1.994,"Home load":1.374},{"Hour":"3 am","Export":0,"Grid":1.999,"Solar use":0,"SoC":95.58333333333333,"Charge":0.547,"Discharge":-0.001,"Solar":0,"Battery grid":0.547,"Home load":1.452},{"Hour":"4 am","Export":0,"Grid":1.423,"Solar use":0,"SoC":100,"Charge":0,"Discharge":-0.016,"Solar":0,"Battery grid":0,"Home load":1.439},{"Hour":"5 am","Export":0,"Grid":1.48,"Solar use":0,"SoC":99.3076923076923,"Charge":0,"Discharge":-0.015,"Solar":0,"Battery grid":0,"Home load":1.495},{"Hour":"6 am","Export":0,"Grid":1.578,"Solar use":0.009000000000000001,"SoC":98.91666666666667,"Charge":0,"Discharge":-0.008,"Solar":0.009000000000000001,"Battery grid":0,"Home load":1.596},{"Hour":"7 am","Export":-0.032,"Grid":0.113,"Solar use":0.047,"SoC":84.08333333333333,"Charge":0,"Discharge":-1.7690000000000001,"Solar":0.079,"Battery grid":0,"Home load":1.93},{"Hour":"8 am","Export":-0.066,"Grid":0.084,"Solar use":0.241,"SoC":54.07692307692308,"Charge":0,"Discharge":-1.479,"Solar":0.308,"Battery grid":0,"Home load":1.805},{"Hour":"9 am","Export":-0.002,"Grid":0.09,"Solar use":0.609,"SoC":29.416666666666668,"Charge":0,"Discharge":-0.9490000000000001,"Solar":0.612,"Battery grid":0,"Home load":1.649},{"Hour":"10 am","Export":-0.009000000000000001,"Grid":0.28700000000000003,"Solar use":0.882,"SoC":20.75,"Charge":0.28,"Discharge":-0.314,"Solar":0.892,"Battery grid":0.28,"Home load":1.203},{"Hour":"11 am","Export":-0.052000000000000005,"Grid":0.049,"Solar use":1.131,"SoC":22.076923076923077,"Charge":0.261,"Discharge":-0.125,"Solar":1.183,"Battery grid":0.17500000000000002,"Home load":1.044},{"Hour":"12 pm","Export":-0.026000000000000002,"Grid":0.06,"Solar use":1.678,"SoC":32,"Charge":0.888,"Discharge":0,"Solar":1.705,"Battery grid":0.06,"Home load":0.851},{"Hour":"1 pm","Export":-0.024,"Grid":0,"Solar use":1.788,"SoC":47.583333333333336,"Charge":1.004,"Discharge":0,"Solar":1.813,"Battery grid":0,"Home load":0.784},{"Hour":"2 pm","Export":-0.02,"Grid":0.105,"Solar use":2.176,"SoC":70,"Charge":1.493,"Discharge":0,"Solar":2.196,"Battery grid":0.105,"Home load":0.787},{"Hour":"3 pm","Export":-0.355,"Grid":0.3,"Solar use":1.407,"SoC":94.33333333333333,"Charge":0.88,"Discharge":0,"Solar":1.762,"Battery grid":0.3,"Home load":0.8270000000000001},{"Hour":"4 pm","Export":-0.201,"Grid":0.23500000000000001,"Solar use":0.59,"SoC":100,"Charge":0,"Discharge":-0.016,"Solar":0.792,"Battery grid":0,"Home load":0.842},{"Hour":"5 pm","Export":-0.005,"Grid":0.008,"Solar use":0.165,"SoC":91,"Charge":0,"Discharge":-1.103,"Solar":0.171,"Battery grid":0,"Home load":1.2770000000000001},{"Hour":"6 pm","Export":-0.108,"Grid":0.991,"Solar use":0,"SoC":66.92307692307692,"Charge":0,"Discharge":-1.808,"Solar":0.011,"Battery grid":0,"Home load":2.702},{"Hour":"7 pm","Export":-0.001,"Grid":0.015,"Solar use":0,"SoC":38.8,"Charge":0,"Discharge":-1.5,"Solar":0,"Battery grid":0,"Home load":1.514},{"Hour":"8 pm","Export":null,"Grid":null,"Solar use":null,"SoC":null,"Charge":null,"Discharge":null,"Solar":null,"Battery grid":null,"Home load":null},{"Hour":"9 pm","Export":null,"Grid":null,"Solar use":null,"SoC":null,"Charge":null,"Discharge":null,"Solar":null,"Battery grid":null,"Home load":null},{"Hour":"10 pm","Export":null,"Grid":null,"Solar use":null,"SoC":null,"Charge":null,"Discharge":null,"Solar":null,"Battery grid":null,"Home load":null},{"Hour":"11 pm","Export":null,"Grid":null,"Solar use":null,"SoC":null,"Charge":null,"Discharge":null,"Solar":null,"Battery grid":null,"Home load":null}]`

	daydata, err := daydata.UnmarshalDayData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Len(t, daydata, 24, "Day data must have 24 items")

	hour0 := daydata[0]
	assert.Equal(t, "12 am", hour0.Hour)
	assert.Equal(t, 0.0, *hour0.Export)
	assert.Equal(t, 1.326, *hour0.Grid)
	assert.Equal(t, 0.0, *hour0.SolarUse)
	assert.Equal(t, 20.0, *hour0.SoC)
	assert.Equal(t, 0.0, *hour0.Charge)
	assert.Equal(t, 0.0, *hour0.Discharge)
	assert.Equal(t, 0.0, *hour0.Solar)
	assert.Equal(t, 0.0, *hour0.BatteryGrid)
	assert.Equal(t, 1.326, *hour0.HomeLoad)

	influx0 := hour0.GetInfluxFields()
	assert.Equal(t, "12 am", influx0["Hour"])
	assert.Equal(t, 0.0, influx0["Export"])
	assert.Equal(t, 1.326, influx0["Grid"])
	assert.Equal(t, 0.0, influx0["SolarUse"])
	assert.Equal(t, 20.0, influx0["SoC"])
	assert.Equal(t, 0.0, influx0["Charge"])
	assert.Equal(t, 0.0, influx0["Discharge"])
	assert.Equal(t, 0.0, influx0["Solar"])
	assert.Equal(t, 0.0, influx0["BatteryGrid"])
	assert.Equal(t, 1.326, influx0["HomeLoad"])

	hour22 := daydata[22]
	assert.Equal(t, "10 pm", hour22.Hour)
	assert.Nil(t, hour22.Export)
	assert.Nil(t, hour22.Grid)
	assert.Nil(t, hour22.SolarUse)
	assert.Nil(t, hour22.SoC)
	assert.Nil(t, hour22.Charge)
	assert.Nil(t, hour22.Discharge)
	assert.Nil(t, hour22.Solar)
	assert.Nil(t, hour22.BatteryGrid)
	assert.Nil(t, hour22.HomeLoad)

}

// {"deviceStatus":1,"dPowerFlow":5,"export":0,"import":1026,"load":1026,"solar":0,"soc":19,"charge":0,"gridPowerOutage":0,"temperature":23.1}
//[{"Day":1,"Solar use":7.1899999999999995,"Grid":16.6,"Export":-2.31,"Solar":9.5,"Battery grid":6.3,"Home load":23.490000000000002},{"Day":2,"Solar use":2.6700000000000004,"Grid":22.93,"Export":-0.03,"Solar":2.7,"Battery grid":9.2,"Home load":25.2},{"Day":3,"Solar use":8.49,"Grid":12.64,"Export":-0.91,"Solar":9.4,"Battery grid":7.6,"Home load":20.83},{"Day":4,"Solar use":8.89,"Grid":24.88,"Export":-2.11,"Solar":11,"Battery grid":6.8,"Home load":33.37},{"Day":5,"Solar use":5.5,"Grid":33.15,"Export":-0.1,"Solar":5.6,"Battery grid":10.1,"Home load":38.35},{"Day":6,"Solar use":9.49,"Grid":37.63,"Export":-0.61,"Solar":10.1,"Battery grid":9.5,"Home load":46.82000000000001},{"Day":7,"Solar use":9.26,"Grid":36.17,"Export":-0.14,"Solar":9.4,"Battery grid":9.6,"Home load":44.83},{"Day":8,"Solar use":10.81,"Grid":16.8,"Export":-0.69,"Solar":11.5,"Battery grid":9.2,"Home load":27.41},{"Day":9,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":10,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":11,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":12,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":13,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":14,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":15,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":16,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":17,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":18,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":19,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":20,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":21,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":22,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":23,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":24,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":25,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":26,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":27,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":28,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":29,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null},{"Day":30,"Solar use":null,"Grid":0,"Export":0,"Solar":null,"Battery grid":null,"Home load":null}]
//[{"Month":"Jan","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null},{"Month":"Feb","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null},{"Month":"Mar","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null},{"Month":"Apr","Solar use":44.489999999999995,"Grid":396.15,"Export":-25.11,"Battery grid":0.6,"Home load":440.73999999999995},{"Month":"May","Solar use":0,"Grid":702.1,"Export":-4.68,"Battery grid":65.2,"Home load":698.32},{"Month":"Jun","Solar use":115.74,"Grid":914.49,"Export":-2.36,"Battery grid":279.4,"Home load":1014.5299999999999},{"Month":"Jul","Solar use":113.51,"Grid":1251.37,"Export":-1.69,"Battery grid":309.3,"Home load":1351.08},{"Month":"Aug","Solar use":198.75,"Grid":925.29,"Export":-9.55,"Battery grid":274.6,"Home load":1110.54},{"Month":"Sep","Solar use":62.300000000000004,"Grid":201.91,"Export":-6.9,"Battery grid":68.3,"Home load":261.41},{"Month":"Oct","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null},{"Month":"Nov","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null},{"Month":"Dec","Solar use":null,"Grid":0,"Export":0,"Battery grid":null,"Home load":null}]
