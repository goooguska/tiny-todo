package task

import (
	"tiny-todo/internal/dto/task"
	"tiny-todo/internal/model"
)

func (s *Service) UpdateTask(id int, input *task.UpdateInput) error {
	task := &model.Task{
		Id:          uint64(id),
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
	}

	err := s.r.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}
