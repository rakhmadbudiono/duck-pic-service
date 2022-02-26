package config

type ServerConfig struct {
	Host string `envconfig:"SERVER_HOST"`
	Port string `envconfig:"SERVER_PORT" default:"8000"`
}
