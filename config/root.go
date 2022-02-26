package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var root Config

// Config wraps all config types
type Config struct {
	Server serverConfig
}

// New load and return config object
func New() *Config {
	if err := envconfig.Process("SERVER", &root.Server); err != nil {
		log.Fatal(err.Error())
	}

	return &root
}
