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

	HasWriteError() bool
}

type influxDBWriterImpl struct {
	config     *jsontypes.Configuration
	client     influxdb2.Client
	writeAPI   api.WriteAPI
	writeError bool
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

	// iw.deleteData()
	return nil
}

func (iw *influxDBWriterImpl) WriteData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing to InfluxDB")

	currentData := scrape.Data()

	influxFields := iw.getInfluxFields(currentData)
	var stamp = time.Unix(0, currentData.EnergyFlow.LastUpdate*int64(time.Millisecond))
	influxFields["Received"] = fmt.Sprint(stamp)

	Logger.Debug().Msgf("Write to influx Current %s", fmt.Sprint(stamp))
	iw.writeAPI.WritePoint(influxdb2.NewPoint("solar2", nil, influxFields, stamp))

	iw.writeAPI.Flush()
	Logger.Info().Msg("Done Writing to InfluxDB")
}

func (iw *influxDBWriterImpl) getInfluxFields(currentData jsontypes.DataResponseData) map[string]interface{} {
	m := make(map[string]interface{})

	m["current-load"] = currentData.EnergyFlow.Home
	m["current-solar"] = currentData.EnergyFlow.Solar

	if currentData.EnergyFlow.GridImport {
		m["current-import"] = currentData.EnergyFlow.Grid
		m["current-export"] = 0.0
	} else if currentData.EnergyFlow.GridExport {
		m["current-export"] = currentData.EnergyFlow.Grid
		m["current-import"] = 0.0
	} else {
		m["current-import"] = 0.0
		m["current-export"] = 0.0
	}
	if currentData.EnergyFlow.BatteryUsed {
		m["current-battery-use"] = currentData.EnergyFlow.Battery
		m["current-battery-charge"] = 0.0
	} else if currentData.EnergyFlow.BatteryCharged {
		m["current-battery-charge"] = currentData.EnergyFlow.Battery
		m["current-battery-use"] = 0.0
	} else {
		m["current-battery-charge"] = 0.0
		m["current-battery-use"] = 0.0
	}

	m["current-grid-import"] = currentData.EnergyFlow.GridImport
	m["current-grid-export"] = currentData.EnergyFlow.GridExport
	m["current-battery-used-value"] = currentData.EnergyFlow.BatteryUsed
	m["current-battery-charged"] = currentData.EnergyFlow.BatteryCharged

	m["flows-threshold"] = currentData.EnergyFlow.Flows.Threshold
	m["flows-solartohome"] = currentData.EnergyFlow.Flows.Solartohome
	m["flows-solartobattery"] = currentData.EnergyFlow.Flows.Solartobattery
	m["flows-solartogrid"] = currentData.EnergyFlow.Flows.Solartogrid
	m["flows-gridtohome"] = currentData.EnergyFlow.Flows.Gridtohome
	m["flows-batterytohome"] = currentData.EnergyFlow.Flows.Batterytohome
	m["flows-batterytogrid"] = currentData.EnergyFlow.Flows.Batterytogrid
	m["flows-gridtobattery"] = currentData.EnergyFlow.Flows.Gridtobattery

	m["battery-capacity"] = currentData.Monitor.Battery.Capacity
	m["battery-charged"] = currentData.Monitor.Battery.Charged
	m["carbon-value"] = currentData.Monitor.Carbon.Value

	m["total-home-usage"] = currentData.Cards.HomeUsage.Value
	m["total-solar-utilization"] = currentData.Cards.SolarUtilization.Value
	m["total-home-usage-total"] = currentData.Cards.HomeUsageTotal.Value
	m["total-solar-util-total"] = currentData.Cards.SolarUtilTotal.Value
	m["total-grid-import-total"] = currentData.Cards.GridImportTotal.Value
	m["total-grid-export-total"] = currentData.Cards.GridExportTotal.Value

	return m
}

func (iw *influxDBWriterImpl) deleteData() {
	// Get DeleteAPI
	deleteAPI := iw.client.DeleteAPI()

	// Define the time range for the delete operation
	// If you want to delete all data, use a very large time range
	start := time.Unix(0, 0) // Start time (Unix epoch)
	end := time.Now().UTC()  // End time (current time)

	// Perform the delete operation
	err := deleteAPI.DeleteWithName(context.Background(), iw.config.InfluxDB.Org, iw.config.InfluxDB.Bucket, start, end, "")
	if err != nil {
		fmt.Printf("Error deleting data: %v\n", err)
	} else {
		fmt.Println("Data deleted successfully.")
	}
}
