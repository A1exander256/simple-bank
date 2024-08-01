package main

import (
	"context"
	"os"

	"github.com/A1exander256/simple-bank/cmd"
	"github.com/A1exander256/simple-bank/internal/config"
	"github.com/rs/zerolog"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logLevel, err := cfg.LogLevel()
	if err != nil {
		panic(err)
	}

	log := zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Caller().Logger()
	ctx := log.WithContext(context.Background())

	log.Info().Msg("the application is launching")

	exitCode := 0

	if err := cmd.Run(ctx, cfg); err != nil {
		log.Err(err).Send()

		exitCode = 1
	}

	os.Exit(exitCode)
}
