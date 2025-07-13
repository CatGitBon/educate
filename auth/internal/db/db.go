package db

import (
	"database/sql"
	"fmt"

	"github.com/vctrl/currency-service/auth/internal/config"
)

func NewDatabaseConnection(cfg config.DatabaseConfig) (*sql.DB, string, error) {

	dsn := cfg.ToDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, "", fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, dsn, nil
}
