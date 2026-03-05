package task

func (s *Service) DeleteTask(id int) error {
	if err := s.r.DeleteTask(id); err != nil {
		return err
	}

	return nil
}
