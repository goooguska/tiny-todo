package task

import "tiny-todo/internal/model"

func (s *Service) GetAll() ([]model.Task, error) {
	tasks, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Service) GetById(id int) (*model.Task, error) {
	task, err := s.r.GetById(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
