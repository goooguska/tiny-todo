package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const dsn = "user=%s password=%s host=%s port=%s dbname=%s sslmode=%s"

type Postgres struct {
	DB *sql.DB
}

type PostgresConfig interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetName() string
	GetSslMode() string
}

func New(cfg PostgresConfig) (*Postgres, error) {
	dsnStr := fmt.Sprintf(dsn, cfg.GetUser(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetName(), cfg.GetSslMode())

	db, err := connect(dsnStr)
	if err != nil {
		return &Postgres{}, err
	}

	return &Postgres{DB: db}, nil
}

func connect(dsn string) (*sql.DB, error) {
	return sql.Open("pgx", dsn)
}
