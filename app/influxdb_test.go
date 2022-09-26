package app

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	protocol "github.com/influxdata/line-protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
	currentdata "github.com/mwinters-stuff/solar-zero-scrape-golang/app/currentdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/daydata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/monthdata"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/yeardata"
	mockapp "github.com/mwinters-stuff/solar-zero-scrape-golang/internal/mocks/app"
	mocks "github.com/mwinters-stuff/solar-zero-scrape-golang/internal/mocks/influxdb2"
	"github.com/stretchr/testify/suite"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *LogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type InfluxDBWriterTestSuite struct {
	suite.Suite
	writer       *InfluxDBWriter
	mockWriteAPI *mocks.WriteAPI
	loghook      LogHook
	errCh        chan error
	currentData  currentdata.CurrentData
	json         map[string]string
}

func (*InfluxDBWriterTestSuite) makeConfig() config.Configuration {
	config := config.Configuration{}
	config.InfluxDB.HostURL = "https://influxdb.url/"
	config.InfluxDB.Token = "ANTOKENTHATSBIG"
	config.InfluxDB.Bucket = "solarzero/autogen"
	config.InfluxDB.Org = "example.org"
	return config
}

func (suite *InfluxDBWriterTestSuite) SetupTest() {
	suite.loghook = LogHook{}
	Logger = log.Hook(&suite.loghook)

	suite.currentData = currentdata.CurrentData{
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

	suite.json = map[string]string{
		"current": "",
		"day": `[
			{ "Hour": "12 am", "Export": 0, "Grid": 1.326, "Solar use": 0, "SoC": 20, "Charge": 0, "Discharge": 0, "Solar": 0, "Battery grid": 0, "Home load": 1.326 },
			{ "Hour": "1 am", "Export": 0, "Grid": 3.4410000000000003, "Solar use": 0, "SoC": 35.916666666666664, "Charge": 1.995, "Discharge": 0, "Solar": 0, "Battery grid": 1.995, "Home load": 1.445 },
			{ "Hour": "12 pm", "Export": -0.026000000000000002, "Grid": 0.06, "Solar use": 1.678, "SoC": 32, "Charge": 0.888, "Discharge": 0, "Solar": 1.705, "Battery grid": 0.06, "Home load": 0.851 },
			{ "Hour": "5 pm", "Export": -0.005, "Grid": 0.008, "Solar use": 0.165, "SoC": 91, "Charge": 0, "Discharge": -1.103, "Solar": 0.171, "Battery grid": 0, "Home load": 1.2770000000000001 },
			{ "Hour": "10 pm", "Export": null, "Grid": null, "Solar use": null, "SoC": null, "Charge": null, "Discharge": null, "Solar": null, "Battery grid": null, "Home load": null }
		]`,
		"month": `[
			{ "Day": 1, "Solar use": 7.19, "Grid": 16.6, "Export": -2.31, "Solar": 9.5, "Battery grid": 6.3, "Home load": 23.5 },
			{ "Day": 10, "Solar use": null, "Grid": 0, "Export": 0, "Solar": null, "Battery grid": null, "Home load": null }
		]`,
		"year": `[
			{ "Month": "Jan", "Solar use": null, "Grid": 0, "Export": 0, "Battery grid": null, "Home load": null },
			{ "Month": "Jun", "Solar use": 115.74, "Grid": 914.49, "Export": -2.36, "Battery grid": 279.4, "Home load": 1014.52 }
		]`,
	}

	suite.ConnectInflux()
}

func String(s string) *string {
	return &s
}

func (suite *InfluxDBWriterTestSuite) ConnectInflux() {
	configuration := suite.makeConfig()

	mockClient := mocks.NewClient(suite.T())

	health := domain.HealthCheck{
		Checks:  nil,
		Commit:  String("CommitTag"),
		Message: String("A Message"),
		Name:    "A Name",
		Status:  "Status",
		Version: String("1.0.0"),
	}

	mockClient.EXPECT().Health(mock.AnythingOfType("*context.emptyCtx")).Return(&health, nil).Once()

	suite.mockWriteAPI = mocks.NewWriteAPI(suite.T())
	mockClient.EXPECT().WriteAPI("example.org", "solarzero/autogen").Return(suite.mockWriteAPI).Once()

	suite.errCh = make(chan error)
	suite.mockWriteAPI.EXPECT().Errors().Return(suite.errCh).Once()

	suite.writer = NewInfluxDBWriter(&configuration)

	suite.writer.Connect(mockClient)
}

func (suite *InfluxDBWriterTestSuite) TestConnect() {

	assert.Eventually(suite.T(), func() bool {

		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "InfluxDB Health: A Message Status 1.0.0 "
	}, time.Second, time.Millisecond*100)

}

func (suite *InfluxDBWriterTestSuite) TestErrorLog() {

	suite.errCh <- errors.New("Failed")

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "InfluxDB Write error: Failed"
	}, time.Second, time.Millisecond*100)

	assert.True(suite.T(), suite.writer.WriteError)

}

func (suite *InfluxDBWriterTestSuite) TestWriteCurrentData() {

	mockSZS := mockapp.NewSolarZeroScrape(suite.T())

	mockSZS.EXPECT().CurrentData().Return(suite.currentData).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar", point.Name())
			assert.Nil(suite.T(), point.TagList())

			values := [...]*protocol.Field{
				{Key: "Charge", Value: int64(0)},
				{Key: "DPowerFlow", Value: int64(5)},
				{Key: "DeviceStatus", Value: int64(1)},
				{Key: "Export", Value: int64(0)},
				{Key: "GridPowerOutage", Value: int64(0)},
				{Key: "Import", Value: int64(1026)},
				{Key: "Load", Value: int64(1026)},
				{Key: "Soc", Value: int64(19)},
				{Key: "Solar", Value: int64(0)},
				{Key: "Temperature", Value: float64(23.1)},
			}
			assert.ElementsMatch(suite.T(), values, point.FieldList())

		}).Once()

	suite.writer.writeCurrentData(mockSZS)

}

func (suite *InfluxDBWriterTestSuite) TestWriteDayData() {

	mockSZS := mockapp.NewSolarZeroScrape(suite.T())

	daydata, err := daydata.UnmarshalDayData([]byte(suite.json["day"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().DayData().Return(daydata).Once()

	t := time.Now()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
			hourStr := fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{
					Key:   "date",
					Value: hourStr,
				},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(0)},
				{Key: "Charge", Value: float64(0)},
				{Key: "Discharge", Value: float64(0)},
				{Key: "Export", Value: float64(0)},
				{Key: "Grid", Value: float64(1.326)},
				{Key: "HomeLoad", Value: float64(1.326)},
				{Key: "Hour", Value: hourStr},
				{Key: "SoC", Value: float64(20)},
				{Key: "Solar", Value: float64(0)},
				{Key: "SolarUse", Value: float64(0)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())

		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
			hourStr := fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 1, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{Key: "date", Value: hourStr},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(1.995)},
				{Key: "Charge", Value: float64(1.995)},
				{Key: "Discharge", Value: float64(0)},
				{Key: "Export", Value: float64(0)},
				{Key: "Grid", Value: float64(3.4410000000000003)},
				{Key: "HomeLoad", Value: float64(1.445)},
				{Key: "Hour", Value: hourStr},
				{Key: "SoC", Value: float64(35.916666666666664)},
				{Key: "Solar", Value: float64(0)},
				{Key: "SolarUse", Value: float64(0)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())
		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
			hourStr := fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{Key: "date", Value: hourStr},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(0.06)},
				{Key: "Charge", Value: float64(0.888)},
				{Key: "Discharge", Value: float64(0)},
				{Key: "Export", Value: float64(-0.026000000000000002)},
				{Key: "Grid", Value: float64(0.06)},
				{Key: "HomeLoad", Value: float64(0.851)},
				{Key: "Hour", Value: hourStr},
				{Key: "SoC", Value: float64(32)},
				{Key: "Solar", Value: float64(1.705)},
				{Key: "SolarUse", Value: float64(1.678)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())
		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
			hourStr := fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{Key: "date", Value: hourStr},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(0)},
				{Key: "Charge", Value: float64(0)},
				{Key: "Discharge", Value: float64(-1.103)},
				{Key: "Export", Value: float64(-0.005)},
				{Key: "Grid", Value: float64(0.008)},
				{Key: "HomeLoad", Value: float64(1.2770000000000001)},
				{Key: "Hour", Value: hourStr},
				{Key: "SoC", Value: float64(91)},
				{Key: "Solar", Value: float64(0.171)},
				{Key: "SolarUse", Value: float64(0.165)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())
		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
			tags := [...]*protocol.Tag{
				{Key: "date", Value: fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 22, 0, 0, 0, time.Local))}}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
		}).Times(0)

	suite.writer.writeDayData(mockSZS)

}

func (suite *InfluxDBWriterTestSuite) TestWriteMonthData() {

	mockSZS := mockapp.NewSolarZeroScrape(suite.T())

	monthdata, err := monthdata.UnmarshalMonthData([]byte(suite.json["month"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().MonthData().Return(monthdata).Once()

	t := time.Now()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-month", point.Name())
			dayStr := fmt.Sprint(time.Date(t.Year(), t.Month(), int(1), 0, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{
					Key:   "date",
					Value: dayStr,
				},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(6.3)},
				{Key: "Day", Value: int64(1)},
				{Key: "Export", Value: float64(-2.31)},
				{Key: "Grid", Value: float64(16.6)},
				{Key: "HomeLoad", Value: float64(23.5)},
				{Key: "Solar", Value: float64(9.5)},
				{Key: "SolarUse", Value: float64(7.19)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())

		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-month", point.Name())
			tags := [...]*protocol.Tag{
				{Key: "date", Value: fmt.Sprint(time.Date(t.Year(), t.Month(), int(10), 0, 0, 0, 0, time.Local))}}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
		}).Times(0)

	suite.writer.writeMonthData(mockSZS)

}

func (suite *InfluxDBWriterTestSuite) TestWriteYearData() {

	mockSZS := mockapp.NewSolarZeroScrape(suite.T())

	yeardata, err := yeardata.UnmarshalYearData([]byte(suite.json["year"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().YearData().Return(yeardata).Once()

	t := time.Now()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-year", point.Name())
			dayStr := fmt.Sprint(time.Date(t.Year(), time.Month(6), int(1), 0, 0, 0, 0, time.Local))
			tags := [...]*protocol.Tag{
				{
					Key:   "date",
					Value: dayStr,
				},
			}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
			values := [...]*protocol.Field{
				{Key: "BatteryGrid", Value: float64(279.4)},
				{Key: "Month", Value: "Jun"},
				{Key: "Export", Value: float64(-2.36)},
				{Key: "Grid", Value: float64(914.49)},
				{Key: "HomeLoad", Value: float64(1014.52)},
				{Key: "SolarUse", Value: float64(115.74)},
			}

			assert.ElementsMatch(suite.T(), values, point.FieldList())

		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-year", point.Name())
			tags := [...]*protocol.Tag{
				{Key: "date", Value: fmt.Sprint(time.Date(t.Year(), time.Month(1), int(1), 0, 0, 0, 0, time.Local))}}

			assert.ElementsMatch(suite.T(), tags, point.TagList())
		}).Times(0)

	suite.writer.writeYearData(mockSZS)

}

func (suite *InfluxDBWriterTestSuite) TestWriteData() {

	mockSZS := mockapp.NewSolarZeroScrape(suite.T())

	mockSZS.EXPECT().CurrentData().Return(suite.currentData).Once()

	daydata, err := daydata.UnmarshalDayData([]byte(suite.json["day"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().DayData().Return(daydata).Once()

	monthdata, err := monthdata.UnmarshalMonthData([]byte(suite.json["month"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().MonthData().Return(monthdata).Once()

	yeardata, err := yeardata.UnmarshalYearData([]byte(suite.json["year"]))
	assert.Nil(suite.T(), err, "Err is not nil")

	mockSZS.EXPECT().YearData().Return(yeardata).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar", point.Name())
		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-day", point.Name())
		}).Times(4)

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-month", point.Name())
		}).Once()

	suite.mockWriteAPI.EXPECT().WritePoint(mock.Anything).
		Run(func(point *write.Point) {
			assert.Equal(suite.T(), "solar-year", point.Name())
		}).Once()

	suite.mockWriteAPI.EXPECT().Flush().Once()

	suite.writer.WriteData(mockSZS)

}

func TestInfluxDBWriterSuite(t *testing.T) {
	suite.Run(t, new(InfluxDBWriterTestSuite))
}
