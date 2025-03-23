package solarzero

import (
	"fmt"
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
	PublishHomeAssistantDiscovery()
}

type mqttClientImpl struct {
	config          *jsontypes.Configuration
	client          mqtt.Client
	baseSensorTopic string
	baseBoolTopic   string
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

	// mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	// mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	// mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	// mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
	mq.baseSensorTopic = fmt.Sprintf("homeassistant/sensor/%[1]s/%[1]s", mq.config.Mqtt.BaseTopic)
	mq.baseBoolTopic = fmt.Sprintf("homeassistant/binary_sensor/%[1]s/%[1]s", mq.config.Mqtt.BaseTopic)

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

	mq.WriteCurrentData(scrape)

	Logger.Info().Msg("Done Writing to MQTT")
}

func (mq *mqttClientImpl) publish(topic string, payload string) {
	Logger.Debug().Msgf("MQTT %s -> %s", topic, payload)
	t := mq.client.Publish(fmt.Sprintf("%s/%s", mq.config.Mqtt.BaseTopic, topic), 0, false, payload)
	go func() {
		_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			Logger.Error().Err(t.Error()) // Use your preferred logging technique (or just fmt.Printf)
		}
	}()
}

func formatFloat(value float64) string {
	// return strconv.FormatFloat(value, 'f', 2, 64)
	return strconv.FormatInt(int64(value*1000), 10)
}

func formatFloatN(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64)
}

func formatInt(value int64) string {
	return strconv.FormatInt(value, 10)
}

func (mq *mqttClientImpl) WriteCurrentData(scrape SolarZeroScrape) {
	Logger.Debug().Msgf("Write to mqtt Current %s", fmt.Sprint(time.Now()))
	currentData := scrape.Data()

	mq.publish("current/received", fmt.Sprint(time.Unix(0, currentData.EnergyFlow.LastUpdate*int64(time.Millisecond))))

	mq.publish("current/load", formatFloat(currentData.EnergyFlow.Home))
	mq.publish("current/solar", formatFloat(currentData.EnergyFlow.Solar))

	if currentData.EnergyFlow.GridImport {
		mq.publish("current/import", formatFloat(currentData.EnergyFlow.Grid))
		mq.publish("current/export", formatFloat(0.0))
	} else if currentData.EnergyFlow.GridExport {
		mq.publish("current/export", formatFloat(currentData.EnergyFlow.Grid))
		mq.publish("current/import", formatFloat(0.0))
	} else {
		mq.publish("current/import", formatFloat(0.0))
		mq.publish("current/export", formatFloat(0.0))
	}
	if currentData.EnergyFlow.BatteryUsed {
		mq.publish("current/battery-use", formatFloat(currentData.EnergyFlow.Battery))
		mq.publish("current/battery-charge", formatFloat(0.0))
	} else if currentData.EnergyFlow.BatteryCharged {
		mq.publish("current/battery-charge", formatFloat(currentData.EnergyFlow.Battery))
		mq.publish("current/battery-use", formatFloat(0.0))
	} else {
		mq.publish("current/battery-charge", formatFloat(0.0))
		mq.publish("current/battery-use", formatFloat(0.0))
	}

	mq.publish("current/grid-import", strconv.FormatBool(currentData.EnergyFlow.GridImport))
	mq.publish("current/grid-export", strconv.FormatBool(currentData.EnergyFlow.GridExport))
	mq.publish("current/battery-used-value", strconv.FormatBool(currentData.EnergyFlow.BatteryUsed))
	mq.publish("current/battery-charged", strconv.FormatBool(currentData.EnergyFlow.BatteryCharged))

	mq.publish("flows/threshold", formatInt(currentData.EnergyFlow.Flows.Threshold))
	mq.publish("flows/solartohome", formatFloat(currentData.EnergyFlow.Flows.Solartohome))
	mq.publish("flows/solartobattery", formatFloat(currentData.EnergyFlow.Flows.Solartobattery))
	mq.publish("flows/solartogrid", formatFloat(currentData.EnergyFlow.Flows.Solartogrid))
	mq.publish("flows/gridtohome", formatFloat(currentData.EnergyFlow.Flows.Gridtohome))
	mq.publish("flows/batterytohome", formatFloat(currentData.EnergyFlow.Flows.Batterytohome))
	mq.publish("flows/batterytogrid", formatFloat(currentData.EnergyFlow.Flows.Batterytogrid))
	mq.publish("flows/gridtobattery", formatFloat(currentData.EnergyFlow.Flows.Gridtobattery))

	mq.publish("battery/capacity", formatFloat(currentData.Monitor.Battery.Capacity))
	mq.publish("battery/charged", formatFloatN(currentData.Monitor.Battery.Charged))
	mq.publish("carbon/value", formatFloatN(currentData.Monitor.Carbon.Value))
	// mq.publish("carbon/desc", currentData.Monitor.Carbon.Desc)

	// mq.publish("home/comments", currentData.Monitor.Home.Comments)
	// mq.publish("home/value1", formatInt(currentData.Monitor.Home.Value1.Value))
	// mq.publish("home/value2", formatInt(currentData.Monitor.Home.Value2.Value))

	// mq.publish("solar/comments", currentData.Monitor.Solar.Comments)
	// mq.publish("solar/value1", formatInt(currentData.Monitor.Solar.Value1.Value))
	// mq.publish("solar/value2", formatInt(currentData.Monitor.Solar.Value2.Value))

	mq.publish("total/home-usage", formatInt(currentData.Cards.HomeUsage.Value))
	mq.publish("total/solar-utilization", formatInt(currentData.Cards.SolarUtilization.Value))
	mq.publish("total/home-usage-total", formatFloat(currentData.Cards.HomeUsageTotal.Value))
	mq.publish("total/solar-util-total", formatFloat(currentData.Cards.SolarUtilTotal.Value))
	mq.publish("total/grid-import-total", formatFloat(currentData.Cards.GridImportTotal.Value))
	mq.publish("total/grid-export-total", formatFloat(currentData.Cards.GridExportTotal.Value))

	customer := scrape.Customer()
	today := time.Now()
	currentPrice := 0.0
	currentGridState := currentData.Tou.Grid.State
	if today.Weekday() == time.Saturday || today.Weekday() == time.Sunday {
		if currentGridState == "shoulder" {
			currentPrice = customer.Provider.Details.Weekends.Shoulder.Rate
		}
		if currentGridState == "offPeak" {
			currentPrice = customer.Provider.Details.Weekends.OffPeak.Rate
		}
		if currentGridState == "peak" {
			currentPrice = customer.Provider.Details.Weekends.Peak.Rate
		}
	} else {
		if currentGridState == "shoulder" {
			currentPrice = customer.Provider.Details.Weekdays.Shoulder.Rate
		}
		if currentGridState == "offPeak" {
			currentPrice = customer.Provider.Details.Weekdays.OffPeak.Rate
		}
		if currentGridState == "peak" {
			currentPrice = customer.Provider.Details.Weekdays.Peak.Rate
		}
	}
	mq.publish("power-price/current", formatFloatN(currentPrice))
}

func (mq *mqttClientImpl) publishDiscoveryTopic(topic string, payload string) {
	Logger.Debug().Msgf("MQTT %s -> %s", topic, payload)

	t := mq.client.Publish(topic, 0, true, payload)
	go func() {
		_ = t.Done() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			Logger.Error().Err(t.Error()) // Use your preferred logging technique (or just fmt.Printf)
		}
	}()
}

func (mq *mqttClientImpl) publishDiscovery(group, what, label, unit_of_meas, dev_class, measurement, icon string) {
	mq.publishDiscoveryTopic(fmt.Sprintf("%s-%s-%s/config", mq.baseSensorTopic, group, what),
		fmt.Sprintf(
			`
    {
      "unique_id": "%[1]s-%[2]s-%[3]s",
      "name": "%[4]s",
      "state_topic": "%[1]s/%[2]s/%[3]s",
      "unit_of_meas": "%[5]s",
      "suggested_display_precision": 0,
      "device_class": "%[6]s",
      "state_class": "%[7]s",
      "icon": "%[8]s",
      "device": {
        "suggested_area": "Outside",
        "ids": "%[1]s",
        "name": "Solar Zero"
      },
      "availability": {
        "topic": "%[1]s/status",
        "payload_available": "ONLINE",
        "payload_not_available": "OFFLINE"
      }
    }`,
			mq.config.Mqtt.BaseTopic, // 1
			group,                    // 2
			what,                     // 3
			label,                    // 4
			unit_of_meas,             // 5
			dev_class,                // 6
			measurement,              // 7
			icon,                     // 8
		))
}

func (mq *mqttClientImpl) publishDiscoveryNoIcon(group, what, label, unit_of_meas, dev_class, measurement string) {
	mq.publishDiscoveryTopic(fmt.Sprintf("%s-%s-%s/config", mq.baseSensorTopic, group, what),
		fmt.Sprintf(
			`
    {
      "unique_id": "%[1]s-%[2]s-%[3]s",
      "name": "%[4]s",
      "state_topic": "%[1]s/%[2]s/%[3]s",
      "unit_of_meas": "%[5]s",
      "suggested_display_precision": 0,
      "device_class": "%[6]s",
      "state_class": "%[7]s",
      "device": {
        "suggested_area": "Outside",
        "ids": "%[1]s",
        "name": "Solar Zero"
      },
      "availability": {
        "topic": "%[1]s/status",
        "payload_available": "ONLINE",
        "payload_not_available": "OFFLINE"
      }
    }`,
			mq.config.Mqtt.BaseTopic, // 1
			group,                    // 2
			what,                     // 3
			label,                    // 4
			unit_of_meas,             // 5
			dev_class,                // 6
			measurement,              // 7
		))
}

func (mq *mqttClientImpl) publishBoolDiscovery(group, what, label, dev_class, icon, payload_on, payload_off string) {
	mq.publishDiscoveryTopic(fmt.Sprintf("%s-%s-%s/config", mq.baseBoolTopic, group, what),
		fmt.Sprintf(
			`
    {
      "unique_id": "%[1]s-%[2]s-%[3]s",
      "name": "%[4]s",
      "state_topic": "%[1]s/%[2]s/%[3]s",
      "device_class": "%[5]s",
      "icon": "%[6]s",
      "payload_on": "%[7]s",
			"payload_off": "%[8]s",
			"state_color": true,
      "device": {
        "suggested_area": "Outside",
        "ids": "%[1]s",
        "name": "Solar Zero"
      },
      "availability": {
        "topic": "%[1]s/status",
        "payload_available": "ONLINE",
        "payload_not_available": "OFFLINE"
      }
    }`,
			mq.config.Mqtt.BaseTopic, // 1
			group,                    // 2
			what,                     // 3
			label,                    // 4
			dev_class,                // 5
			icon,                     // 6
			payload_on,               // 7
			payload_off,              // 8
		))
}

func (mq *mqttClientImpl) publishDiscovery2DP(group, what, label, unit_of_meas, dev_class, measurement, icon string) {
	mq.publishDiscoveryTopic(fmt.Sprintf("%s-%s-%s/config", mq.baseSensorTopic, group, what),
		fmt.Sprintf(
			`
    {
      "unique_id": "%[1]s-%[2]s-%[3]s",
      "name": "%[4]s",
      "state_topic": "%[1]s/%[2]s/%[3]s",
      "unit_of_meas": "%[5]s",
      "suggested_display_precision": 2,
      "device_class": "%[6]s",
      "state_class": "%[7]s",
      "icon": "%[8]s",
      "device": {
        "suggested_area": "Outside",
        "ids": "%[1]s",
        "name": "Solar Zero"
      },
      "availability": {
        "topic": "%[1]s/status",
        "payload_available": "ONLINE",
        "payload_not_available": "OFFLINE"
      }
    }`,
			mq.config.Mqtt.BaseTopic, // 1
			group,                    // 2
			what,                     // 3
			label,                    // 4
			unit_of_meas,             // 5
			dev_class,                // 6
			measurement,              // 7
			icon,                     // 8
		))
}

func (mq *mqttClientImpl) publishDiscoveryLastResetMidnight(group, what, label, unit_of_meas, dev_class, measurement, icon string) {
	mq.publishDiscoveryTopic(fmt.Sprintf("%s-%s-%s/config", mq.baseSensorTopic, group, what),
		fmt.Sprintf(
			`
    {
      "unique_id": "%[1]s-%[2]s-%[3]s",
      "name": "%[4]s",
      "state_topic": "%[1]s/%[2]s/%[3]s",
      "unit_of_meas": "%[5]s",
      "suggested_display_precision": 0,
      "device_class": "%[6]s",
      "state_class": "%[7]s",
      "icon": "%[8]s",
      "last_reset_value_template": "{{ now().replace(hour=0, minute=0, second=0, microsecond=0).isoformat() }}",
      "device": {
        "suggested_area": "Outside",
        "ids": "%[1]s",
        "name": "Solar Zero"
      },
      "availability": {
        "topic": "%[1]s/status",
        "payload_available": "ONLINE",
        "payload_not_available": "OFFLINE"
      }
    }`,
			mq.config.Mqtt.BaseTopic, // 1
			group,                    // 2
			what,                     // 3
			label,                    // 4
			unit_of_meas,             // 5
			dev_class,                // 6
			measurement,              // 7
			icon,                     // 8
		))
}

func (mq *mqttClientImpl) PublishHomeAssistantDiscovery() {

	mq.publishDiscovery("current", "load", "House Load", "W", "power", "measurement", "mdi:home-lightning-bolt")
	mq.publishDiscovery("current", "solar", "Solar", "W", "power", "measurement", "mdi:solar-power")
	mq.publishDiscovery("current", "import", "Grid Import", "W", "power", "measurement", "mdi:home-import-outline")
	mq.publishDiscovery("current", "export", "Grid Export", "W", "power", "measurement", "mdi:home-export-outline")

	mq.publishDiscoveryLastResetMidnight("current", "battery-use", "Battery Use", "Wh", "energy", "total", "mdi:battery-arrow-down")
	mq.publishDiscoveryLastResetMidnight("current", "battery-charge", "Battery Charge", "Wh", "energy", "total", "mdi:battery-charging-80")

	mq.publishDiscovery("total", "home-usage", "Home Usage", "%", "energy", "measurement", "mdi:home-lightning-bolt-outline")
	mq.publishDiscovery("total", "solar-utilization", "Solar Utilization", "%", "energy", "measurement", "mdi:solar-power")

	mq.publishDiscoveryLastResetMidnight("total", "home-usage-total", "Home Usage Total", "Wh", "energy", "total", "mdi:home-lightning-bolt")
	mq.publishDiscoveryLastResetMidnight("total", "solar-util-total", "Solar Util Total", "Wh", "energy", "total", "mdi:solar-power-variant")
	mq.publishDiscoveryLastResetMidnight("total", "grid-import-total", "Grid Import Total", "Wh", "energy", "total", "mdi:transmission-tower-import")
	mq.publishDiscoveryLastResetMidnight("total", "grid-export-total", "Grid Export Total", "Wh", "energy", "total", "mdi:transmission-tower-export")

	mq.publishDiscovery("battery", "capacity", "Battery Capacity", "Wh", "energy", "total_increasing", "mdi:home-battery-outline")
	mq.publishDiscoveryNoIcon("battery", "charged", "Battery SOC", "%", "battery", "measurement")

	mq.publishDiscovery2DP("power-price", "current", "Current Grid Rate", "NZD/kWh", "monetary", "measurement", "mdi:currency-usd")

	mq.publishBoolDiscovery("current", "grid-import", "Importing From Grid", "power", "mdi:transmission-tower-import", "true", "false")
	mq.publishBoolDiscovery("current", "grid-export", "Exporting To Grid", "power", "mdi:transmission-tower-export", "true", "false")

	mq.publishBoolDiscovery("current", "battery-used-value", "Using Battery", "battery_charging", "mdi:battery-charging-80", "true", "false")
	mq.publishBoolDiscovery("current", "battery-charged", "Charging Battery", "battery_charging", "mdi:battery-charging-10", "true", "false")

	mq.publishDiscovery("flows", "solartohome", "Solar To Home", "Wh", "energy", "measurement", "mdi:home-export-outline")
	mq.publishDiscovery("flows", "solartobattery", "Solar To Battery", "Wh", "energy", "measurement", "mdi:battery-charging-80")
	mq.publishDiscovery("flows", "solartogrid", "Solar To Grid", "Wh", "energy", "measurement", "mdi:transmission-tower-export")
	mq.publishDiscovery("flows", "gridtohome", "Grid To Home", "Wh", "energy", "measurement", "mdi:transmission-tower-import")
	mq.publishDiscovery("flows", "batterytohome", "Battery To Home", "Wh", "energy", "measurement", "mdi:battery-arrow-down")
	mq.publishDiscovery("flows", "batterytogrid", "Battery To Grid", "Wh", "energy", "measurement", "mdi:battery-arrow-up")
	mq.publishDiscovery("flows", "gridtobattery", "Grid To Battery", "Wh", "energy", "measurement", "mdi:transmission-tower-import")

	// mq.publishDiscovery("carbon", "value", "Carbon Usage", "ppm", "co2", "measurement", "mdi:molecule-co2")

}
