package migrations

import (
	"errors"

	"github.com/golang-migrate/migrate"
)

func RunPgMigrations(dsn string) error {
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	path := "file://auth/internal/migrations"

	m, err := migrate.New(
		path,
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
