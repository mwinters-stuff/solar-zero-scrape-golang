package solarzero

type SolarZeroOptions struct {
	Config string `short:"c" long:"config" description:"Solar Zero Configuration"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`

	Username string `long:"solarzero-username" description:"Solar Zero Username"`
	Password string `long:"solarzero-password" description:"Solar Zero Password"`
}

type MQTTOptions struct {
	ServerURL string `long:"mqtt-server-url" description:"MQTT Server URL" `
	Topic     string `long:"mqtt-topic" description:"MQTT Topic" `
	Username  string `long:"mqtt-username" description:"MQTT Username" `
	Password  string `long:"mqtt-password" description:"MQTT Password" `
}

type AllSolarZeroOptions struct {
	SolarZeroOptions SolarZeroOptions
	MQTTOptions      MQTTOptions
}
