package task

import (
	"fmt"
)

func (r *repository) DeleteTask(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TableName)
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed delete task %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get row affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("failed delete task with id %d", id)
	}

	return nil
}
