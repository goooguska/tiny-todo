package app

import (
	"database/sql"
	"log/slog"
	"os"
	"tiny-todo/internal/config"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *sql.DB

	Repositories *Repositories
	Services     *Services
}

func Bootstrap(db *sql.DB, config *config.Config) *App {
	app := &App{
		DB:     db,
		Logger: initLogger(),
		Config: config,
	}

	app.Repositories = NewRepositories(app)
	app.Services = NewServices(app)

	return app
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
