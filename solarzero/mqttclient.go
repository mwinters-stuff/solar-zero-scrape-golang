package solarzero

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

var (
	NewMQTTClient = NewMQTTClientImpl
)

type MQTTClient interface {
	Connect() error

	WriteData(scrape SolarZeroScrape)
	WriteCurrentData(scrape SolarZeroScrape)
	WriteDayData(scrape SolarZeroScrape)
	WriteMonthData(scrape SolarZeroScrape)
	WriteYearData(scrape SolarZeroScrape)
}

type mqttClientImpl struct {
	config *jsontypes.Configuration
	client mqtt.Client
}

func NewMQTTClientImpl(config *jsontypes.Configuration) MQTTClient {
	s := &mqttClientImpl{
		config: config,
	}
	return s
}

var defaultPublushHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	Logger.Info().Msgf("TOPIC: %s\n", msg.Topic())
	Logger.Info().Msgf("MSG: %s\n", msg.Payload())
}

func (mq *mqttClientImpl) Connect() error {

	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	opts := mqtt.NewClientOptions().
		AddBroker(mq.config.Mqtt.URL).
		SetClientID("solar-zero-scrape").
		SetUsername(mq.config.Mqtt.Username).
		SetPassword(mq.config.Mqtt.Password).
		SetWill(fmt.Sprintf("%s/%s", mq.config.Mqtt.BaseTopic, "status"), "OFFLINE", 0, true).
		SetAutoReconnect(true).
		SetDefaultPublishHandler(defaultPublushHandler).
		SetOnConnectHandler(mq.OnConnectHandler)

	mq.client = mqtt.NewClient(opts)
	if token := mq.client.Connect(); token.Wait() && token.Error() != nil {
		Logger.Fatal().Err(token.Error())
	}

	return nil
}

func (mq *mqttClientImpl) OnConnectHandler(client mqtt.Client) {
	mq.publish("status", "ONLINE")
	mq.PublishHomeAssistantDiscovery()
}

func (mq *mqttClientImpl) WriteData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing to MQTT")
	
  mq.publish("status", "ONLINE")
	mq.PublishHomeAssistantDiscovery()

  mq.WriteCurrentData(scrape)
	mq.WriteDayData(scrape)
	mq.WriteMonthData(scrape)
	mq.WriteYearData(scrape)

	Logger.Info().Msg("Done Writing to MQTT")
}

func (mq *mqttClientImpl) publish(topic string, payload string) {
	t := mq.client.Publish(fmt.Sprintf("%s/%s", mq.config.Mqtt.BaseTopic, topic), 0, true, payload)
	go func() {
		_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			Logger.Error().Err(t.Error()) // Use your preferred logging technique (or just fmt.Printf)
		}
	}()
}

func (mq *mqttClientImpl) WriteCurrentData(scrape SolarZeroScrape) {
	Logger.Debug().Msgf("Write to mqtt Current %s", fmt.Sprint(time.Now()))
	currentData := scrape.CurrentData()
	if currentData.DeviceStatus == 1 {
		fields := currentData.GetMQTTFields()

		stamp, _ := parseLocalTimestamp(currentData.ReceivedDate)
		mq.publish("current/received", fmt.Sprint(stamp))

		for key, value := range fields {
			mq.publish(fmt.Sprintf("current/%s", key), value)
		}
	}

	// record totals

	mq.publish("today/used", strconv.FormatInt(int64((scrape.ElectricityUse().ElectricityUse)*1000.0), 10))
	mq.publish("today/exported", strconv.FormatInt(int64((scrape.SolarUse().ExportAmount)*1000.0), 10))
	mq.publish("today/imported", strconv.FormatInt(int64((scrape.SolarVsGrid().GridAmount)*1000.0), 10))
	mq.publish("today/solar-used", strconv.FormatInt(int64((scrape.SolarUse().SelfUseAmount)*1000.0), 10))
	mq.publish("today/solar-used-percent", strconv.FormatInt(scrape.SolarUse().ExportPercent, 10))
	mq.publish("today/exported-percent", strconv.FormatInt(scrape.SolarUse().SelfUsePercent, 10))

	mq.publish("today/solar-vs-grid-grid-percent", strconv.FormatInt(scrape.SolarVsGrid().GridPercent, 10))
	mq.publish("today/solar-vs-grid-grid", strconv.FormatInt(int64((scrape.SolarVsGrid().GridAmount)*1000.0), 10))
	mq.publish("today/solar-vs-grid-solar-percent", strconv.FormatInt(scrape.SolarVsGrid().SolarPercent, 10))
	mq.publish("today/solar-vs-grid-solar", strconv.FormatInt(int64((scrape.SolarVsGrid().SolarAmount)*1000.0), 10))

	mq.publish("today/solar", strconv.FormatInt(int64((scrape.SolarUse().ExportAmount+scrape.SolarUse().SelfUseAmount)*1000.0), 10))
}

func (mq *mqttClientImpl) WriteDayData(scrape SolarZeroScrape) {
	// for _, hourData := range scrape.DayData() {

	// 	mqttFields := hourData.GetMQTTFields()
	// 	if mqttFields != nil {
	// 		stamp, _ := parseLocalTimestamp(hourData.ReceivedDate)

	// 		for key, value := range *mqttFields {
	// 			mq.publish(fmt.Sprintf("day/%d/%s", stamp.Hour(), key), value)
	// 			// Logger.Debug().Msgf("Write to mqtt Hour %d", stamp.Hour())
	// 		}
	// 	}

	// }

}

func (mq *mqttClientImpl) WriteMonthData(scrape SolarZeroScrape) {
	// for _, dayData := range scrape.MonthData() {

	// 	mqttFields := dayData.GetMQTTFields()
	// 	if mqttFields != nil {
	// 		stamp, _ := parseLocalTimestamp(dayData.ReceivedDate)

	// 		for key, value := range *mqttFields {
	// 			mq.publish(fmt.Sprintf("month/%d/%s", stamp.Day(), key), value)
	// 		}
	// 	}

	// }
}

func (mq *mqttClientImpl) WriteYearData(scrape SolarZeroScrape) {
	// for _, monthData := range scrape.YearData() {

	// 	mqttFields := monthData.GetMQTTFields()
	// 	if mqttFields != nil {
	// 		stamp, _ := parseLocalTimestamp(monthData.ReceivedDate)

	// 		for key, value := range *mqttFields {
	// 			mq.publish(fmt.Sprintf("year/%d/%s", stamp.Year(), key), value)
	// 		}
	// 	}

	// }
}

func (mq *mqttClientImpl) publishTopic(topic string, payload string) {
	t := mq.client.Publish(topic, 0, true, payload)
	go func() {
		_ = t.Done() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			Logger.Error().Err(t.Error()) // Use your preferred logging technique (or just fmt.Printf)
		}
	}()
}

func (mq *mqttClientImpl) PublishHomeAssistantDiscovery() {
	baseTopic := fmt.Sprintf("homeassistant/sensor/%[1]s/%[1]s", mq.config.Mqtt.BaseTopic)

	mq.publishTopic(fmt.Sprintf("%s-temperature/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-temperature",
    "name": "Temperature",
    "stat_t": "%[1]s/current/temperature",
    "unit_of_meas": "°C",
    "sug_dsp_prc": 1,
    "dev_cla": "temperature",
    "icon": "mdi:thermometer",
    "dev": {
      "sa": "Outside",
      "ids": "%[1]s",
      "name": "Solar Zero"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	// baseTopic = fmt.Sprintf("homeassistant/power/%[1]s/%[1]s-", mq.config.Mqtt.BaseTopic)

	mq.publishTopic(fmt.Sprintf("%s-load/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-load",
    "name": "Load",
    "stat_t": "%[1]s/current/load",
    "unit_of_meas": "W",
    "sug_dsp_prc": 1,
    "dev_cla": "power",
    "state_class": "measurement",
    "icon": "mdi:power-plug-outline",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-solar/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-solar",
    "name": "Solar",
    "stat_t": "%[1]s/current/solar",
    "unit_of_meas": "W",
    "sug_dsp_prc": 1,
    "dev_cla": "power",
		"state_class": "measurement",
    "icon": "mdi:solar-power",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-import/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-import",
    "name": "Import",
    "stat_t": "%[1]s/current/import",
    "unit_of_meas": "W",
    "sug_dsp_prc": 1,
    "dev_cla": "power",
		"state_class": "measurement",
    "icon": "mdi:transmission-tower-import",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))
	mq.publishTopic(fmt.Sprintf("%s-export/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-export",
    "name": "Export",
    "stat_t": "%[1]s/current/export",
    "unit_of_meas": "W",
    "sug_dsp_prc": 1,
    "dev_cla": "power",
		"state_class": "measurement",
    "icon": "mdi:transmission-tower-export",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))
	mq.publishTopic(fmt.Sprintf("%s-soc/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-soc",
    "name": "Battery SOC",
    "stat_t": "%[1]s/current/soc",
    "unit_of_meas": "%%",
    "sug_dsp_prc": 1,
    "dev_cla": "battery",
		"state_class": "measurement",
    "icon": "mdi:battery-high",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-batteryvoltage/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-batteryvoltage",
    "name": "Battery Volts",
    "stat_t": "%[1]s/current/batteryvoltage",
    "unit_of_meas": "V",
    "sug_dsp_prc": 1,
    "dev_cla": "voltage",
		"state_class": "measurement",
    "icon": "mdi:battery-plus-outline",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-batterycurrent/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-batterycurrent",
    "name": "Battery Current",
    "stat_t": "%[1]s/current/batterycurrent",
    "unit_of_meas": "A",
    "sug_dsp_prc": 1,
    "dev_cla": "current",
		"state_class": "measurement",
    "icon": "mdi:battery-charging-medium",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-discharge/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-discharge",
    "name": "Battery Discharge",
    "stat_t": "%[1]s/current/discharge",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
		"state_class": "total",
    "icon": "mdi:battery-20",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-charge/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-charge",
    "name": "Battery Charge",
    "stat_t": "%[1]s/current/charge",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
		"state_class": "total",
    "icon": "mdi:battery-charging-90",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-used/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-used",
    "name": "Used Today",
    "stat_t": "%[1]s/today/used",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
    "state_class": "total_increasing",
    "icon": "mdi:home-lightning-bolt",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-solar/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-solar",
    "name": "Solar Generated Today",
    "stat_t": "%[1]s/today/solar",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
    "state_class": "total_increasing",
    "icon": "mdi:solar-panel-large",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-exported/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-exported",
    "name": "Exported Today",
    "stat_t": "%[1]s/today/exported",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
    "state_class": "total_increasing",
    "icon": "mdi:transmission-tower-export",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-imported/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-imported",
    "name": "Imported Today",
    "stat_t": "%[1]s/today/imported",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
    "state_class": "total_increasing",
    "icon": "mdi:transmission-tower-export",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-solar-used/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-solar-used",
    "name": "Solar Used Today",
    "stat_t": "%[1]s/today/solar-used",
    "unit_of_meas": "Wh",
    "sug_dsp_prc": 1,
    "dev_cla": "energy",
    "state_class": "total_increasing",
    "icon": "mdi:solar-power-variant-outline",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-solar-used-percent/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-solar-used-percent",
    "name": "%% Solar Used Today",
    "stat_t": "%[1]s/today/solar-used-percent",
    "unit_of_meas": "%%",
    "sug_dsp_prc": 1,
    "state_class": "measurement",
    "icon": "mdi:sun-angle-outline",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))

	mq.publishTopic(fmt.Sprintf("%s-today-exported-percent/config", baseTopic), fmt.Sprintf(`
  {
    "uniq_id": "%[1]s-today-exported-percent",
    "name": "%% Solar Exported Today",
    "stat_t": "%[1]s/today/exported-percent",
    "unit_of_meas": "%%",
    "sug_dsp_prc": 1,
    "state_class": "measurement",
    "icon": "mdi:transmission-tower-export",
    "dev": {
      "ids": "%[1]s"
    },
    "avty": {
      "t": "%[1]s/status",
      "pl_avail": "ONLINE",
      "pl_not_avail": "OFFLINE"
    }
  }
	`, mq.config.Mqtt.BaseTopic))
}
