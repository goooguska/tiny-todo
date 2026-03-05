package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tiny-todo/internal/http/handler/task"
	repo "tiny-todo/internal/repository/sqlite/task"
	service "tiny-todo/internal/service/task"

	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = setupLogger()

	db, err := setupDB()
	if err != nil {
		slog.Error("setup db failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	s := service.NewService(repo.NewRepository(db))
	server := http.Server{
		Addr:    ":8080",
		Handler: setupRoutes(s),
	}

	slog.Info("starting server", "addr", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("start server failed", "error", err)
	}
}

func setupRoutes(s *service.Service) http.Handler {
	mux := http.NewServeMux()
	v := validator.New()
	h := task.NewHandler(s, v)

	mux.HandleFunc("GET /api/v1/tasks", h.GetAll)
	mux.HandleFunc("GET /api/v1/tasks/{id}", h.GetById)
	mux.HandleFunc("POST /api/v1/tasks", h.CreateTask)

	return mux
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

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
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
