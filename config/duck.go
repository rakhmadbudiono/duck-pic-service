package config

type duckConfig struct {
	SpaceFolder string `envconfig:"DUCK_SPACE_FOLDER"`
	CountDuck   int64  `envconfig:"DUCK_COUNT_DUCK"`
}
