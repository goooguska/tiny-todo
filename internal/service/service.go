package service

import (
	"tiny-todo/internal/dto/task"
	"tiny-todo/internal/model"
)

type TaskService interface {
	CreateTask(input *task.CreateInput) error
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
	GetById(id int) (*model.Task, error)
	GetAll() ([]model.Task, error)
}
