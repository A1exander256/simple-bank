package build

import (
	"context"

	"github.com/A1exander256/simple-bank/internal/config"
)

type Builder struct {
	config   *config.Config
	shutdown shutdown
}

func New(_ context.Context, cfg *config.Config) *Builder {
	return &Builder{config: cfg} //nolint:exhaustruct
}
