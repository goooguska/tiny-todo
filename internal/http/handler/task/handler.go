package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tiny-todo/internal/dto/task"
	service "tiny-todo/internal/service/task"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	s *service.Service
	v *validator.Validate
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input task.CreateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.v.Struct(&input)
	if err != nil {
		http.Error(w, "unprocessable request body", http.StatusUnprocessableEntity)
		return
	}

	if err := h.s.CreateTask(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	task, err := h.s.GetById(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.s.GetAll()
	if err != nil {
		http.Error(w, "tasks not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func NewHandler(s *service.Service, v *validator.Validate) *Handler {
	return &Handler{s: s, v: v}
}
