package service

import (
	"tiny-todo/internal/dto/task"
	"tiny-todo/internal/model"
)

type TaskService interface {
	CreateTask(input *task.CreateInput) error
	UpdateTask(id string, input *task.UpdateInput) error
	DeleteTask(id string) error
	GetById(id string) (*model.Task, error)
	GetAll() ([]model.Task, error)
}
