package repository

import (
	"tiny-todo/internal/dto/task"
	"tiny-todo/internal/model"
)

type TaskRepository interface {
	CreateTask(input *task.CreateInput) error
	UpdateTask(id string, input *task.UpdateInput) error
	DeleteTask(id string) error
	GetById(id string) (*model.Task, error)
	GetAll() ([]model.Task, error)
}
