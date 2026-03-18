package handler

import "tiny-todo/internal/http/handler/task"

type Handlers struct {
	TaskHandler *task.Handler
}
