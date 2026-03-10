package app

import (
	"tiny-todo/internal/repository"
	repo "tiny-todo/internal/repository/sqlite/task"
)

type Repositories struct {
	TaskRepository repository.TaskRepository
}

func NewRepositories(a *App) *Repositories {
	return &Repositories{
		TaskRepository: repo.NewRepository(a.DB),
	}
}
