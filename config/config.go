package config

import "github.com/caarlos0/env"

type Config struct {
	RestServerPort string `env:"REST_SERVER_PORT,required"`
	PostgresURL    string `env:"POSTGRES_URL"`
	LogLevel       string `env:"LOG_LEVEL"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
