package solarzero

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

var (
	NewInfluxDBWriter = NewInfluxDBWriterImpl
)

type InfluxDBWriter interface {
	Connect(client influxdb2.Client) error

	WriteData(scrape SolarZeroScrape)
	WriteDailyData(scrape SolarZeroScrape)
	WriteCurrentData(scrape SolarZeroScrape)
	WriteDayData(scrape SolarZeroScrape)
	WriteMonthData(scrape SolarZeroScrape)
	WriteYearData(scrape SolarZeroScrape)

	HasWriteError() bool
}

type influxDBWriterImpl struct {
	config     *jsontypes.Configuration
	client     influxdb2.Client
	writeAPI   api.WriteAPI
	writeError bool
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

func NewInfluxDBWriterImpl(config *jsontypes.Configuration) InfluxDBWriter {
	s := &influxDBWriterImpl{
		config: config,
	}
	return s
}

func (iw *influxDBWriterImpl) HasWriteError() bool {
	return iw.writeError
}

func (iw *influxDBWriterImpl) Connect(client influxdb2.Client) error {
	iw.writeError = false
	iw.client = client
	health, err := iw.client.Health(context.Background())

	if err != nil {
		Logger.Fatal().Err(err)
	}

	Logger.Info().Msgf("InfluxDB Health: %s %s %s ", *health.Message, health.Status, *health.Version)
	iw.writeAPI = iw.client.WriteAPI(iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket)

	errorsCh := iw.writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			Logger.Error().Msgf("InfluxDB Write error: %s", err.Error())
			iw.writeError = true
		}
	}()

	return nil
}

func (iw *influxDBWriterImpl) WriteData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing to InfluxDB")
	iw.WriteCurrentData(scrape)
	iw.WriteDayData(scrape)
	iw.writeAPI.Flush()
	Logger.Info().Msg("Done Writing to InfluxDB")
}

func (iw *influxDBWriterImpl) WriteDailyData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing Daily to InfluxDB")
	iw.WriteMonthData(scrape)
	iw.WriteYearData(scrape)
	iw.writeAPI.Flush()
	Logger.Info().Msg("Done Writing Daily to InfluxDB")
}

func (iw *influxDBWriterImpl) WriteCurrentData(scrape SolarZeroScrape) {
	currentData := scrape.CurrentData()
	influxFields := currentData.GetInfluxFields()
	stamp, _ := parseLocalTimestamp(currentData.ReceivedDate)
	influxFields["Received"] = fmt.Sprint(stamp)

	Logger.Debug().Msgf("Write to influx Current %s", fmt.Sprint(stamp))
	iw.writeAPI.WritePoint(influxdb2.NewPoint("solar", nil, influxFields, stamp))
}

func (iw *influxDBWriterImpl) WriteDayData(scrape SolarZeroScrape) {

	for _, hourData := range scrape.DayData() {
		influxFields := hourData.GetInfluxFields()
		if influxFields != nil {
			stamp, _ := parseLocalTimestamp(hourData.ReceivedDate)
			(*influxFields)["Hour"] = fmt.Sprint(stamp)
			iw.writeAPI.WritePoint(influxdb2.NewPoint("solar-day",
				map[string]string{
					"date": fmt.Sprint(stamp),
				},
				*influxFields,
				stamp))
			Logger.Debug().Msgf("Write to influx Hour %s", fmt.Sprint(stamp))
		}

	}
}

func (iw *influxDBWriterImpl) WriteMonthData(scrape SolarZeroScrape) {
	for _, dayData := range scrape.MonthData() {
		influxFields := dayData.GetInfluxFields()
		if influxFields != nil {
			stamp, _ := parseLocalTimestamp(dayData.ReceivedDate)

			iw.writeAPI.WritePoint(influxdb2.NewPoint("solar-month",
				map[string]string{
					"date": fmt.Sprint(stamp),
				},
				*influxFields,
				stamp))
			Logger.Debug().Msgf("Write to influx Day %s", fmt.Sprint(stamp))
		}
	}
}

func (iw *influxDBWriterImpl) WriteYearData(scrape SolarZeroScrape) {
	for _, monthData := range scrape.YearData() {
		influxFields := monthData.GetInfluxFields()
		if influxFields != nil {
			stamp, _ := parseLocalTimestamp(monthData.ReceivedDate)

			iw.writeAPI.WritePoint(influxdb2.NewPoint("solar-year",
				map[string]string{
					"date": fmt.Sprint(stamp),
				},
				*influxFields,
				stamp))
			Logger.Debug().Msgf("Write to influx Month %s", fmt.Sprint(stamp))
		}
	}
}
