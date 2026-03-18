package task

import (
	"fmt"
	"strings"
	"tiny-todo/internal/dto/task"
)

func (r *repository) UpdateTask(id string, input *task.UpdateInput) error {
	var setParts []string
	var args []any

	if input.Title != nil {
		args = append(args, *input.Title)
		setParts = append(setParts, fmt.Sprintf("title=$%d", len(args)))
	}
	if input.Description != nil {
		args = append(args, *input.Description)
		setParts = append(setParts, fmt.Sprintf("description=$%d", len(args)))
	}

	if input.Completed != nil {
		args = append(args, *input.Completed)
		setParts = append(setParts, fmt.Sprintf("completed=$%d", len(args)))
	}

	if len(args) == 0 {
		return nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", TableName, strings.Join(setParts, ","), len(args))

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get row affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("task with id %s not found", id)
	}

	return nil
}
