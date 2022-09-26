package main

import (
	"flag"
	"time"

	"github.com/go-co-op/gocron"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configFile := flag.String("config", "", "Config file")
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if *configFile == "" {
		println("Usage: solar-zero-scrape config.json")
		return
	}

	config, err := config.LoadConfiguration(*configFile)
	if err != nil {
		log.Panic().Msg("LoadConfiguration " + err.Error())
	}

	influxdb := app.NewInfluxDBWriter(&config)
	err = influxdb.Connect(influxdb2.NewClient(config.InfluxDB.HostURL, config.InfluxDB.Token))
	if err != nil {
		log.Panic().Msgf("InfluxDB Connect %s", err.Error())
	}

	log.Info().Msg("Authenticating")

	scrape := app.NewSolarZeroScrape(config)

	s := gocron.NewScheduler(time.Local)

	for scrape.AuthenticateFully() {
		s.Every(5).Minutes().Do(func() {
			log.Info().Msgf("Get Data at ", time.Now())
			success := scrape.GetData()
			if success {
				influxdb.WriteData(scrape)
			} else {
				log.Error().Msg("GetData Failed, Reauthenticating")
				s.Stop()
			}
		})
		s.StartBlocking()

	}

	log.Error().Msg("AuthenicateFully Failed, Exiting")

}
