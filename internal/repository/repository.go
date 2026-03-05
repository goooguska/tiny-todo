package repository

import (
	"tiny-todo/internal/model"
)

type TaskRepository interface {
	CreateTask(task *model.Task) error
	DeleteTask(id int) error
	GetById(id int) (*model.Task, error)
	GetAll() ([]model.Task, error)
}
