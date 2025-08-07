package config

import "github.com/caarlos0/env"

type Config struct {
	DummySecret    string `env:"DUMMY_SECRET,required"`
	RestServerPort string `env:"REST_SERVER_PORT,required"`
	LogLevel       string `env:"LOG_LEVEL"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
