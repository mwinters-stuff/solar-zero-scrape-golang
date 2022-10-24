package solarzero

type SolarZeroOptions struct {
	Config string `short:"c" long:"config" description:"Solar Zero Configuration"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`

	Username string `long:"solarzero-username" description:"Solar Zero Username"`
	Password string `long:"solarzero-password" description:"Solar Zero Password"`
}

type InfluxDBOptions struct {
	Token   string `long:"influx-token" description:"Token for influx access" required:"true"`
	HostURL string `long:"influx-host-url" description:"Influx Host URL" required:"true"`
	Org     string `long:"influx-org" description:"Influx Organization" required:"true"`
	Bucket  string `long:"influx-bucket" description:"Influx Bucket" required:"true"`
}

type OtherOptions struct {
	UserPoolId             string `long:"user-pool-id" default:"us-west-2_NoMpv1v1A"`
	ClientId               string `long:"client-id" default:"6mgtqq7vvf7eo3r3qrsg6kl1tf"`
	ApiRegion              string `long:"api-region" default:"us-west-2"`
	ApiGatewayURL          string `long:"api-gateway-url" default:"https://d6nfzye2cb.execute-api.us-west-2.amazonaws.com"`
	ApiKey                 string `long:"api-key" default:"mA0UW2ldUUQBY3e9bZWq9lCeKQUNCZC9oKidvdbb"`
	ApiSolarZeroApiAddress string `long:"api-solar-zero-api-address" default:"solarzero.pnz.technology"`
}
