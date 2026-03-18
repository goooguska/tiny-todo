package task

import (
	"database/sql"
	"errors"
	"fmt"
	"tiny-todo/internal/model"
)

func (r *repository) GetAll() ([]model.Task, error) {
	query := fmt.Sprintf(
		"SELECT id, title, description, completed FROM %s",
		TableName,
	)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("get tasks query: %w", err)
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
		); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}

	return tasks, nil
}

func (r *repository) GetById(id string) (*model.Task, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE id = $1",
		TableName,
	)

	var task model.Task

	err := r.db.QueryRow(query, id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Completed,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, fmt.Errorf("get task by id: %w", err)
	}

	return &task, nil
}
