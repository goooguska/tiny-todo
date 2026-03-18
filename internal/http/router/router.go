package router

import (
	"tiny-todo/internal/http/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handlers *handler.Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Get("/", handlers.TaskHandler.GetAll)
		r.Post("/", handlers.TaskHandler.CreateTask)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.TaskHandler.GetById)
			r.Patch("/", handlers.TaskHandler.UpdateTask)
			r.Delete("/", handlers.TaskHandler.DeleteTask)
		})
	})

	return r
}
