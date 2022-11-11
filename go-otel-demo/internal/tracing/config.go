package tracing

type Config struct {
	ReceiverEndpoint string `mapstructure:"TRACING_RECEIVER_ENDPOINT"`
	Environment      string `mapstructure:"TRACING_ENVIRONMENT"`
	Service          string `mapstructure:"TRACING_SERVICE"`
	AppVersion       string `mapstructure:"TRACING_APP_VERSION"`
}
