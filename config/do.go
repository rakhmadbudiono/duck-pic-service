package config

type doConfig struct {
	SpaceKey    string `envconfig:"DO_SPACE_KEY"`
	SpaceSecret string `envconfig:"DO_SPACE_SECRET"`
	SpaceHost   string `envconfig:"DO_SPACE_HOST"`
	SpaceBucket string `envconfig:"DO_SPACE_BUCKET"`
}
