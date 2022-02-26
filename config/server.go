package config

type serverConfig struct {
	Host string `envconfig:"SERVER_HOST"`
	Port string `envconfig:"SERVER_PORT" default:"8000"`
}
