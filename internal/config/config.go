package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Config struct {
	App      App
	Postgres Postgres
	HTTP     HTTP
}

type App struct {
	Env      string `envconfig:"APP_ENV" default:"local"`
	Name     string `envconfig:"APP_NAME" default:"app"`
	Version  string `envconfig:"APP_VERSION" default:"v0.0."`
	LogLevel string `envconfig:"APP_LOG_LEVEL" default:"debug"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("reading .env file: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("reading environment: %w", err)
	}

	return &cfg, nil
}

func (c *Config) LogLevel() (zerolog.Level, error) {
	lvl, err := zerolog.ParseLevel(c.App.LogLevel)
	if err != nil {
		return 0, errors.Wrapf(err, "loading log level from config value %q", c.App.LogLevel)
	}

	return lvl, nil
}
