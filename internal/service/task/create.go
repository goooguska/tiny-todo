package task

import (
	"tiny-todo/internal/dto/task"
)

func (s *Service) CreateTask(input *task.CreateInput) error {
	err := s.r.CreateTask(input)
	if err != nil {
		return err
	}

	return nil
}
