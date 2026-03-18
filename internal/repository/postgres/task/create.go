package task

import (
	"fmt"
	"tiny-todo/internal/dto/task"

	"github.com/google/uuid"
)

func (r *repository) CreateTask(input *task.CreateInput) error {
	newId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}
	query := fmt.Sprintf("INSERT INTO %s (id, title, description) VALUES ($1, $2, $3) RETURNING id", TableName)

	_, err = r.db.DB.Exec(query, newId.String(), input.Title, input.Description)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	return nil
}
