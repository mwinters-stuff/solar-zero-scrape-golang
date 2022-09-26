package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
	zerolog "github.com/rs/zerolog/log"
)

var (
	Logger = zerolog.Logger
)

func NewInfluxDBWriter(config *config.Configuration) *InfluxDBWriter {
	s := &InfluxDBWriter{
		config: config,
	}
	return s
}

type InfluxDBWriter struct {
	config     *config.Configuration
	client     influxdb2.Client
	writeAPI   api.WriteAPI
	WriteError bool
}

func (iw *InfluxDBWriter) Connect(client influxdb2.Client) error {
	iw.WriteError = false
	iw.client = client
	health, _ := iw.client.Health(context.Background())
	Logger.Info().Msgf("InfluxDB Health: %s %s %s ", *health.Message, health.Status, *health.Version)
	iw.writeAPI = iw.client.WriteAPI(iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket)

	errorsCh := iw.writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			Logger.Error().Msgf("InfluxDB Write error: %s", err.Error())
			iw.WriteError = true
		}
	}()

	return nil
}

func (iw *InfluxDBWriter) WriteData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing to InfluxDB")
	iw.writeCurrentData(scrape)
	iw.writeDayData(scrape)
	iw.writeMonthData(scrape)
	iw.writeYearData(scrape)
	iw.writeAPI.Flush()
	Logger.Info().Msg("Done Writing to InfluxDB")
}

func (iw *InfluxDBWriter) writeCurrentData(scrape SolarZeroScrape) {
	Logger.Debug().Msgf("Write to influx Current %s", fmt.Sprint(time.Now()))
	currentData := scrape.CurrentData()
	iw.writeAPI.WritePoint(influxdb2.NewPoint("solar", nil, currentData.GetInfluxFields(), time.Now()))
}

func (iw *InfluxDBWriter) writeDayData(scrape SolarZeroScrape) {

	for _, hourData := range scrape.DayData() {
		influxFields := hourData.GetInfluxFields()
		if influxFields != nil {
			hourstr := hourData.Hour
			if hourstr == "12 am" {
				hourstr = "0 am"
			}
			hoursplit := strings.Split(hourstr, " ")
			if len(hoursplit) == 2 {
				hour, _ := strconv.Atoi(hoursplit[0])
				if hoursplit[1] == "pm" && hour != 12 {
					hour += 12
				}
				t := time.Now()
				stamp := time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, time.Local)
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
}

func (iw *InfluxDBWriter) writeMonthData(scrape SolarZeroScrape) {
	for _, dayData := range scrape.MonthData() {
		influxFields := dayData.GetInfluxFields()
		if influxFields != nil {
			t := time.Now()
			stamp := time.Date(t.Year(), t.Month(), int(dayData.Day), 0, 0, 0, 0, time.Local)

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

func (iw *InfluxDBWriter) writeYearData(scrape SolarZeroScrape) {
	for _, monthData := range scrape.YearData() {
		influxFields := monthData.GetInfluxFields()
		if influxFields != nil {
			t := time.Now()
			stamp := time.Date(t.Year(), monthData.GetMonthNum(), 1, 0, 0, 0, 0, time.Local)

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
