package task

func (s *Service) DeleteTask(id string) error {
	if err := s.r.DeleteTask(id); err != nil {
		return err
	}

	return nil
}
