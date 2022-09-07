package app

import (
	"context"
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
	writeAPI api.WriteAPIBlocking
}

func (iw *InfluxDBWriter) Connect() error {
	iw.client = influxdb2.NewClient(iw.config.InfluxDB.HostURL, iw.config.InfluxDB.Token)

	iw.writeAPI = iw.client.WriteAPIBlocking(iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket)
	return nil
}

func (iw *InfluxDBWriter) WriteData(scrape *SolarZeroScrape) error {
	p := influxdb2.NewPoint("solar",
		nil,
		scrape.currentData.GetInfluxFields(),
		time.Now())
	return iw.writeAPI.WritePoint(context.Background(), p)
}
