package solarzero

import (
	"fmt"
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

	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	// mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().
		AddBroker(mq.config.Mqtt.URL).
		SetClientID("solar-zero-scrape").
		SetUsername(mq.config.Mqtt.Username).
		SetPassword(mq.config.Mqtt.Password)

	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(defaultPublushHandler)
	opts.SetPingTimeout(1 * time.Second)

	mq.client = mqtt.NewClient(opts)
	if token := mq.client.Connect(); token.Wait() && token.Error() != nil {
		Logger.Fatal().Err(token.Error())
	}

	return nil
}

func (mq *mqttClientImpl) WriteData(scrape SolarZeroScrape) {
	Logger.Info().Msg("Writing to MQTT")
	mq.WriteCurrentData(scrape)
	mq.WriteDayData(scrape)
	mq.WriteMonthData(scrape)
	mq.WriteYearData(scrape)

	Logger.Info().Msg("Done Writing to MQTT")
}

func (mq *mqttClientImpl) publish(topic string, payload string) {
	t := mq.client.Publish(fmt.Sprintf("%s/%s", mq.config.Mqtt.BaseTopic, topic), 0, true, payload)
	go func() {

		_ = t.Done() // Can also use '<-t.Done()' in releases > 1.2.0
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

}

func (mq *mqttClientImpl) WriteDayData(scrape SolarZeroScrape) {
	for _, hourData := range scrape.DayData() {

		mqttFields := hourData.GetMQTTFields()
		if mqttFields != nil {
			stamp, _ := parseLocalTimestamp(hourData.ReceivedDate)

			for key, value := range *mqttFields {
				mq.publish(fmt.Sprintf("day/%d/%s", stamp.Hour(), key), value)
				Logger.Debug().Msgf("Write to mqtt Hour %d", stamp.Hour())
			}
		}

	}

}

func (mq *mqttClientImpl) WriteMonthData(scrape SolarZeroScrape) {
	for _, dayData := range scrape.MonthData() {

		mqttFields := dayData.GetMQTTFields()
		if mqttFields != nil {
			stamp, _ := parseLocalTimestamp(dayData.ReceivedDate)

			for key, value := range *mqttFields {
				mq.publish(fmt.Sprintf("month/%d/%s", stamp.Day(), key), value)
			}
		}

	}
}

func (mq *mqttClientImpl) WriteYearData(scrape SolarZeroScrape) {
	for _, monthData := range scrape.YearData() {

		mqttFields := monthData.GetMQTTFields()
		if mqttFields != nil {
			stamp, _ := parseLocalTimestamp(monthData.ReceivedDate)

			for key, value := range *mqttFields {
				mq.publish(fmt.Sprintf("year/%d/%s", stamp.Year(), key), value)
			}
		}

	}
}
