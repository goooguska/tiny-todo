package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const dsn = "user=%s password=%s host=%s port=%s dbname=%s sslmode=%s"

type PostgresConfig interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetName() string
	GetSslMode() string
}

func New(cfg PostgresConfig) (*sql.DB, error) {
	dsnStr := fmt.Sprintf(dsn, cfg.GetUser(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetName(), cfg.GetSslMode())

	db, err := connect(dsnStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect db failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db failed: %w", err)
	}

	return db, nil
}
