package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var Root Config

type Config struct {
	Server ServerConfig
}

func New() *Config {
	if err := envconfig.Process("SERVER", &Root.Server); err != nil {
		log.Fatal(err.Error())
	}

	return &Root
}
