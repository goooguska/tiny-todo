package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"tiny-todo/internal/http/handler/task"
	repo "tiny-todo/internal/repository/sqlite/task"
	service "tiny-todo/internal/service/task"

	"github.com/go-playground/validator/v10"
)

type App struct {
	Server *http.Server
	Logger *slog.Logger
}

func New(db *sql.DB) *App {
	l := setupLogger()
	s := service.NewService(repo.NewRepository(db))
	server := &http.Server{
		Addr:    ":8080",
		Handler: setupRoutes(s),
	}

	return &App{Server: server, Logger: l}
}

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}

func setupRoutes(s *service.Service) http.Handler {
	mux := http.NewServeMux()
	v := validator.New()
	h := task.NewHandler(s, v)

	mux.HandleFunc("GET /api/v1/tasks", h.GetAll)
	mux.HandleFunc("GET /api/v1/tasks/{id}", h.GetById)
	mux.HandleFunc("POST /api/v1/tasks", h.CreateTask)
	mux.HandleFunc("DELETE /api/v1/tasks/{id}", h.DeleteTask)
	mux.HandleFunc("PUT /api/v1/tasks/{id}", h.UpdateTask)

	return mux
}
