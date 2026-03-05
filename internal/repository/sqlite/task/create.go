package task

import (
	"fmt"
	"tiny-todo/internal/model"
)

func (r *repository) CreateTask(task *model.Task) error {
	query := fmt.Sprintf("INSERT INTO %s (title, description, completed) VALUES (?, ?, ?)", TableName)

	result, err := r.db.Exec(query, task.Title, task.Description, task.Completed)
	if err != nil {
		return fmt.Errorf("create task: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("get last insert id: %w", err)
	}

	task.Id = uint64(id)

	return nil
}
