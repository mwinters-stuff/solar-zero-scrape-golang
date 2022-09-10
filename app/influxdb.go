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
)

var (
	InfluxDBNewClient = influxdb2.NewClient
)

func NewInfluxDBWriter(config *config.Configuration) *InfluxDBWriter {
	s := &InfluxDBWriter{
		config: config,
	}

	return s
}

type InfluxDBWriter struct {
	config   *config.Configuration
	client   influxdb2.Client
	writeAPI api.WriteAPI
}

func (iw *InfluxDBWriter) Connect() error {
	iw.client = InfluxDBNewClient(iw.config.InfluxDB.HostURL, iw.config.InfluxDB.Token)
	health, _ := iw.client.Health(context.Background())
	println("INFO: InfluxDB Health: ", *health.Message, health.Status, *health.Version)
	iw.writeAPI = iw.client.WriteAPI(iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket)

	errorsCh := iw.writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			fmt.Printf("ERROR: InfluxDB Write error: %s\n", err.Error())
			panic(err)
		}
	}()

	return nil
}

func (iw *InfluxDBWriter) WriteData(scrape *SolarZeroScrape) {
	println("INFO: Writing to InfluxDB")
	iw.writeCurrentData(scrape)
	iw.writeDayData(scrape)
	iw.writeMonthData(scrape)
	iw.writeYearData(scrape)
	iw.writeAPI.Flush()
}

func (iw *InfluxDBWriter) writeCurrentData(scrape *SolarZeroScrape) {
	iw.writeAPI.WritePoint(influxdb2.NewPoint("solar", nil, scrape.currentData.GetInfluxFields(), time.Now()))
}

func (iw *InfluxDBWriter) writeDayData(scrape *SolarZeroScrape) {

	for _, hourData := range scrape.dayData {
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
				fmt.Printf("INFO: Write to influx Hour %s\n", fmt.Sprint(stamp))

			}
		}
	}
}

func (iw *InfluxDBWriter) writeMonthData(scrape *SolarZeroScrape) {
	for _, dayData := range scrape.monthData {
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
			fmt.Printf("INFO: Write to influx Day %s\n", fmt.Sprint(stamp))
		}
	}
}

func (iw *InfluxDBWriter) writeYearData(scrape *SolarZeroScrape) {
	for _, monthData := range scrape.yearData {
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
			fmt.Printf("INFO: Write to influx Month %s\n", fmt.Sprint(stamp))
		}
	}
}
