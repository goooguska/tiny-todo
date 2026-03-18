package task

import (
	"tiny-todo/internal/storage"
)

const TableName = "tasks"

type repository struct {
	db *storage.Postgres
}

func NewRepository(db *storage.Postgres) *repository {
	return &repository{db: db}
}
