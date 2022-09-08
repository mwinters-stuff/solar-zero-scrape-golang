package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
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
	iw.client = influxdb2.NewClient(iw.config.InfluxDB.HostURL, iw.config.InfluxDB.Token)

	iw.writeAPI = iw.client.WriteAPI(iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket)
	return nil
}

func (iw *InfluxDBWriter) WriteData(scrape *SolarZeroScrape) error {
	iw.writeCurrentData(scrape)
	iw.writeDayData(scrape)
	return nil
}

func (iw *InfluxDBWriter) writeCurrentData(scrape *SolarZeroScrape) {
	iw.writeAPI.WritePoint(influxdb2.NewPoint("solar", nil, scrape.currentData.GetInfluxFields(), time.Now()))
}

func (iw *InfluxDBWriter) writeDayData(scrape *SolarZeroScrape) error {

	for _, hourData := range scrape.dayData {
		if hourData.Export != nil {
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
				hourData.Hour = fmt.Sprint(stamp)

				iw.writeAPI.WritePoint(influxdb2.NewPoint("solar-day",
					map[string]string{
						"date": fmt.Sprint(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)),
						"hour": fmt.Sprint(hour),
					},
					hourData.GetInfluxFields(),
					stamp))
			}
		}
	}
	return nil

}

// func (iw *InfluxDBWriter) writeMonthData(scrape *SolarZeroScrape) error {
// 	p := influxdb2.NewPoint("solar-day",
// 		nil,
// 		scrape.currentData.GetInfluxFields(),
// 		time.Now())
// 	return iw.writeAPI.WritePoint(context.Background(), p)
// }

// func (iw *InfluxDBWriter) writeYearData(scrape *SolarZeroScrape) error {
// 	p := influxdb2.NewPoint("solar-day",
// 		nil,
// 		scrape.currentData.GetInfluxFields(),
// 		time.Now())
// 	return iw.writeAPI.WritePoint(context.Background(), p)
// }
