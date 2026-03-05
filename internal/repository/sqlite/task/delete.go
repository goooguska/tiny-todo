package task

import (
	"fmt"
)

func (r *repository) DeleteTask(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TableName)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed delete task %w", err)
	}

	return nil
}
