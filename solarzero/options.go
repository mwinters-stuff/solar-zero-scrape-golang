package solarzero

type SolarZeroOptions struct {
	Config string `short:"c" long:"config" description:"Solar Zero Configuration" required:"true"`
	Debug  bool   `short:"d" long:"debug" description:"Debug Information"`
}
