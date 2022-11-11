package main

import (
	"github.com/spf13/viper"
	"go-otel-demo/internal/config"
	"go-otel-demo/internal/server"
	"log"
)

func main() {
	cfg := ReadConfig(".env")
	server.StartMain(cfg)
}

func ReadConfig(cfgFile string) config.Config {
	var cfg config.Config
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Couldn't read config at", cfgFile, ".", "Error:", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln("Couldn't read config at", cfgFile, ".", "Error:", err)
	}

	return cfg
}
