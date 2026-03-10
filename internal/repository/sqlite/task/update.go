package task

import (
	"fmt"
	"tiny-todo/internal/model"
)

func (r *repository) UpdateTask(task *model.Task) error {
	query := fmt.Sprintf("UPDATE %s SET title=?, description=?, completed=? WHERE id=?", TableName)

	result, err := r.db.Exec(query, task.Title, task.Description, task.Completed, task.Id)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get row affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("task with id %d not found", task.Id)
	}

	return nil
}
