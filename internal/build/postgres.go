package build

import (
	"context"
	"fmt"

	"github.com/A1exander256/simple-bank/internal/migration"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // driver
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
)

func (b *Builder) postgresClient(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %w", err)
	}

	b.shutdown.add(func(_ context.Context) error {
		if err := db.Close(); err != nil {
			return fmt.Errorf("closing postgres connection: %w", err)
		}

		return nil
	})

	return db, nil
}

func (b *Builder) PostgresClient() (*sqlx.DB, error) {
	return b.postgresClient(b.config.Postgres.DSN)
}

func (b *Builder) PostgresClientRO() (*sqlx.DB, error) {
	return b.postgresClient(b.config.Postgres.DSNRO)
}

func (b *Builder) PostgresMigration() (*migrate.Migrate, error) {
	d, err := iofs.New(migration.FS, migration.PostgresPath)
	if err != nil {
		return nil, fmt.Errorf("embed postgres migrations: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, b.config.Postgres.DSN)
	if err != nil {
		return nil, fmt.Errorf("applying postgres migrations: %w", err)
	}

	b.shutdown.add(func(_ context.Context) error {
		return multierr.Append(m.Close()) //nolint:wrapcheck
	})

	return m, nil
}
