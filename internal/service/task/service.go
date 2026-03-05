package task

import (
	"tiny-todo/internal/repository"
)

type Service struct {
	r repository.TaskRepository
}

func NewService(r repository.TaskRepository) *Service {
	return &Service{r: r}
}
