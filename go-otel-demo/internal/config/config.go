package config

import "go-otel-demo/internal/tracing"

type Config struct {
	HttpPort    int            `mapstructure:"HTTP_PORT"`
	RailsAppURL string         `mapstructure:"RAILS_APP_URL"`
	Tracing     tracing.Config `mapstructure:",squash"`
}
