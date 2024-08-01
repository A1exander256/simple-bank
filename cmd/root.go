package cmd

import (
	"context"
	"fmt"

	"github.com/A1exander256/simple-bank/internal/config"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, cfg *config.Config) error {
	//nolint:exhaustruct
	root := &cobra.Command{
		Version: cfg.App.Version,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Usage() //nolint:wrapcheck
		},
	}

	root.AddCommand(
		migrateCmd(ctx, cfg),
		restCmd(ctx, cfg),
	)

	if err := root.ExecuteContext(ctx); err != nil {
		return fmt.Errorf("running application: %w", err)
	}

	return nil
}
