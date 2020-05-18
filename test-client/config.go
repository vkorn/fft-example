package main

import (
	"go-fft-client/shared"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MQTT *shared.MQTTConfig
}

func NewConfig() *Config {
	config := new(Config)
	_, err := flags.Parse(config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return config
}
