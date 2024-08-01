package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/A1exander256/simple-bank/internal/build"
	"github.com/A1exander256/simple-bank/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

func migrateCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	//nolint:exhaustruct
	command := &cobra.Command{
		Use:       "migrate",
		Short:     "run migrations",
		ValidArgs: []string{"postgres"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Usage() //nolint:wrapcheck
		},
	}

	command.AddCommand(postgresCmd(ctx, cfg))

	return command
}

func postgresCmd(ctx context.Context, cfg *config.Config) *cobra.Command {
	//nolint:exhaustruct
	command := &cobra.Command{
		Use:   "postgres",
		Short: "run db migrations for postgres",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Usage() //nolint:wrapcheck
		},
	}

	command.AddCommand(
		up(ctx, cfg, postgres),
		down(ctx, cfg, postgres),
	)

	return command
}

func postgres(ctx context.Context, cfg *config.Config) (*migrate.Migrate, error) {
	b := build.New(ctx, cfg)

	//nolint:wrapcheck
	return b.PostgresMigration()
}

type migrationConstructFn func(context.Context, *config.Config) (*migrate.Migrate, error)

func up(ctx context.Context, cfg *config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "up",
		Short: "up migrations",
		RunE: func(_ *cobra.Command, _ []string) error {
			m, err := constructFn(ctx, cfg)
			if err != nil {
				return fmt.Errorf("constructing migration: %w", err)
			}

			if err = m.Up(); err != nil {
				if errors.Is(err, migrate.ErrNoChange) || errors.Is(err, migrate.ErrNilVersion) {
					return nil
				}

				return fmt.Errorf("up migrations: %w", err)
			}

			return nil
		},
	}
}

func down(ctx context.Context, cfg *config.Config, constructFn migrationConstructFn) *cobra.Command {
	return &cobra.Command{ //nolint:exhaustruct
		Use:   "down",
		Short: "down migrations",
		RunE: func(_ *cobra.Command, _ []string) error {
			m, err := constructFn(ctx, cfg)
			if err != nil {
				return fmt.Errorf("constructing migration: %w", err)
			}

			if err = m.Down(); err != nil {
				return fmt.Errorf("down migrations: %w", err)
			}

			return nil
		},
	}
}
