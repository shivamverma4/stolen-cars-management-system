package config

import (
	"github.com/op/go-logging"

	"stolencarsproject/server/internal/configuration"
)

var config configuration.Config

const namespace = "cars"

func init() {
	config = configuration.Config{
		Namespace:          namespace,
		Deployment:         configuration.DEBUG,
		LogLevel:           logging.INFO,
		LogFilePath:        "../logs/stolencars.log",
		RequestLogFilePath: "../logs/request.log",
		Port:               ApiPort,
	}
}

func GetConfig() *configuration.Config {
	return &config
}
