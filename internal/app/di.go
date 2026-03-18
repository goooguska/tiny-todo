package app

import (
	"fmt"
	"log/slog"
	"os"
	"tiny-todo/internal/config"
	"tiny-todo/internal/http/handler"
	taskHandler "tiny-todo/internal/http/handler/task"
	"tiny-todo/internal/repository"
	taskRepo "tiny-todo/internal/repository/postgres/task"
	"tiny-todo/internal/service"
	taskService "tiny-todo/internal/service/task"
	"tiny-todo/internal/storage"

	"github.com/go-playground/validator/v10"
)

type diContainer struct {
	Handlers *handler.Handlers

	taskRepo    repository.TaskRepository
	taskService service.TaskService
	taskHandler *taskHandler.Handler

	validator *validator.Validate
	cfg       *config.Config
	DB        *storage.Postgres
	logger    *slog.Logger
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) ApiHandlers() *handler.Handlers {
	if d.Handlers == nil {
		d.Handlers = &handler.Handlers{TaskHandler: d.TaskHandler()}
	}

	return d.Handlers
}

func (d *diContainer) TaskHandler() *taskHandler.Handler {
	if d.taskHandler == nil {
		d.taskHandler = taskHandler.NewHandler(d.TaskService(), d.HttpValidator(), d.Logger())
	}
	return d.taskHandler
}

func (d *diContainer) TaskService() service.TaskService {
	if d.taskService == nil {
		d.taskService = taskService.NewService(d.TaskRepo())
	}

	return d.taskService
}

func (d *diContainer) TaskRepo() repository.TaskRepository {
	if d.taskRepo == nil {
		d.taskRepo = taskRepo.NewRepository(d.Db())
	}

	return d.taskRepo
}

func (d *diContainer) HttpValidator() *validator.Validate {
	if d.validator == nil {
		d.validator = validator.New()
	}

	return d.validator
}

func (d *diContainer) Logger() *slog.Logger {
	if d.logger == nil {
		d.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(d.logger)
	}
	return d.logger
}

func (d *diContainer) Db() *storage.Postgres {
	if d.DB == nil {
		db, err := storage.New(&d.Config().DB)
		if err != nil {
			panic(fmt.Sprintf("postgres init failed: %v\n", err))
		}
		err = db.DB.Ping()
		if err != nil {
			panic(fmt.Sprintf("failed to ping postgres: %s\n", err.Error()))
		}

		d.DB = db
	}

	return d.DB
}

func (d *diContainer) Config() *config.Config {
	if d.cfg == nil {
		d.cfg = config.New("config/config.yaml")
	}

	return d.cfg
}
