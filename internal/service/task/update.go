package task

import (
	"tiny-todo/internal/dto/task"
)

func (s *Service) UpdateTask(id string, input *task.UpdateInput) error {
	err := s.r.UpdateTask(id, input)
	if err != nil {
		return err
	}

	return nil
}
