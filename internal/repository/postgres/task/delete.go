package task

import (
	"fmt"
)

func (r *repository) DeleteTask(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", TableName)
	result, err := r.db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed delete task %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get row affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("failed delete task with id %s", id)
	}

	return nil
}
