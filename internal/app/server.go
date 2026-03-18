package app

import (
	"net/http"
	"time"
	"tiny-todo/internal/http/handler/task"

	"github.com/go-playground/validator/v10"
)

func (a *App) HttpServer() *http.Server {
	return &http.Server{
		Addr:        ":" + a.Config.Server.Port,
		ReadTimeout: time.Duration(a.Config.Server.Timeout) * time.Minute,
		Handler:     a.initRoutes(),
	}
}

func (a *App) initRoutes() http.Handler {
	mux := http.NewServeMux()
	v := validator.New()
	h := task.NewHandler(a.Services.TaskService, v, a.Logger)

	mux.HandleFunc("GET /api/v1/tasks", h.GetAll)
	mux.HandleFunc("GET /api/v1/tasks/{id}", h.GetById)
	mux.HandleFunc("POST /api/v1/tasks", h.CreateTask)
	mux.HandleFunc("DELETE /api/v1/tasks/{id}", h.DeleteTask)
	mux.HandleFunc("PUT /api/v1/tasks/{id}", h.UpdateTask)

	return mux
}
