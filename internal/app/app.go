package app

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"tiny-todo/internal/config"
	"tiny-todo/internal/storage"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *sql.DB

	Repositories *Repositories
	Services     *Services
}

func Bootstrap(config *config.Config) *App {
	logger := initLogger()

	db, err := storage.New(&config.DB)
	if err != nil {
		panic(fmt.Sprintf("postgres init failed: %v\n", err))
	}

	app := &App{
		DB:     db,
		Logger: logger,
		Config: config,
	}

	app.Repositories = NewRepositories(db)
	app.Services = NewServices(app)

	return app
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
