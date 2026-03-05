package service

import "tiny-todo/internal/model"

type TaskService interface {
	CreateTask(task *model.Task) error
	GetById(id int) (*model.Task, error)
	GetAll() ([]model.Task, error)
}
