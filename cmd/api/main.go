package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tiny-todo/internal/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := setupDB()
	if err != nil {
		slog.Error("setup db failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	a := app.New(db)

	a.Logger.Info("starting server", "addr", a.Server.Addr)
	if err := a.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.Logger.Error("start server failed", "error", err)
	}
}

func createTasks(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    title TEXT,
		    description TEXT,
		    completed BOOLEAN
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, fmt.Errorf("open db failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db failed: %w", err)
	}

	if err := createTasks(db); err != nil {
		return nil, fmt.Errorf("create tasks failed: %w", err)
	}

	return db, nil
}
