package task

import "database/sql"

const TableName = "tasks"

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}
