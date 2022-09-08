package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/app"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/config"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: solar-zero-scrape config.json")
		return
	}
	argsWithoutProg := os.Args[1:]
	config, err := config.LoadConfiguration(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	if config.DebugLog != nil {
		os.WriteFile(*config.DebugLog, []byte("Debug Log Started\n"), 0644)
	}

	influxdb := app.NewInfluxDBWriter(&config)
	err = influxdb.Connect()
	if err != nil {
		panic(err)
	}

	println("Authenticating ")

	scrape := app.NewSolarZeroScrape(&config)

	for scrape.AuthenticateFully() {
		println("INFO: Getting data at 1 minute interval until data changes")

		// get data once a minute until it changes.
		_, success := scrape.GetData()
		if success {
			influxdb.WriteData(scrape)
		}
		changed := false
		for success && !changed {
			delay := time.NewTimer(1 * time.Minute)
			t := <-delay.C
			fmt.Println("INFO: Get Data at ", t)
			changed, success = scrape.GetData()
		}

		println("INFO: Switching to 5 minute interval")

		for success {
			influxdb.WriteData(scrape)

			delay := time.NewTimer(5 * time.Minute)
			t := <-delay.C
			fmt.Println("INFO: Get Data at ", t)
			_, success = scrape.GetData()
		}

		println("INFO: 5 Minute Interval Finished")
	}

	println("ERROR: Finished")

}
