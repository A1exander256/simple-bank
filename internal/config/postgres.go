package config

type Postgres struct {
	DSN   string `envconfig:"POSTGRES_DSN" default:"postgres://root@localhost:5432/?sslmode=disable"`
	DSNRO string `envconfig:"POSTGRES_DSN_RO" default:"postgres://root@localhost:5432/?sslmode=disable"`
}
