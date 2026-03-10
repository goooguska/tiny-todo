package app

import (
	"tiny-todo/internal/service"
	taskservice "tiny-todo/internal/service/task"
)

type Services struct {
	TaskService service.TaskService
}

func NewServices(a *App) *Services {
	return &Services{
		TaskService: taskservice.NewService(a.Repositories.TaskRepository),
	}
}
