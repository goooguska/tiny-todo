package app

import (
	"database/sql"
	"tiny-todo/internal/repository"
	repo "tiny-todo/internal/repository/postgres/task"
)

type Repositories struct {
	TaskRepository repository.TaskRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		TaskRepository: repo.NewRepository(db),
	}
}
