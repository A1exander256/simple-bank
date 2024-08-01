package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/A1exander256/simple-bank/internal/build"
	"github.com/A1exander256/simple-bank/internal/config"
	"github.com/spf13/cobra"
)

func restCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	//nolint:exhaustruct
	return &cobra.Command{
		Use:   "rest",
		Short: "rest server",
		RunE: func(_ *cobra.Command, _ []string) error {
			builder := build.New(ctx, cfg)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			server, err := builder.RestAPIServer(ctx)
			if err != nil {
				return fmt.Errorf("building rest server: %w", err)
			}

			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("serving rest server: %w", err)
			}

			<-ctx.Done()

			return nil
		},
	}
}
