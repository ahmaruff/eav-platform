package shared

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase(dbPath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
