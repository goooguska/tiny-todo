package task

import (
	"tiny-todo/internal/dto/task"
	"tiny-todo/internal/model"
)

func (s *Service) CreateTask(input *task.CreateInput) error {
	newTask := &model.Task{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	err := s.r.CreateTask(newTask)
	if err != nil {
		return err
	}

	return nil
}
