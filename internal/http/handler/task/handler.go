package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tiny-todo/internal/dto/task"
	httpresponse "tiny-todo/internal/http"
	"tiny-todo/internal/service"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	s service.TaskService
	v *validator.Validate
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input task.CreateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpresponse.SendErrorResponse(w, http.StatusBadRequest, "invalidBody", "Invalid request body", nil)
		return
	}

	err := h.v.Struct(&input)
	if err != nil {
		httpresponse.SendErrorResponse(w, http.StatusUnprocessableEntity, "unprocessableBody", "Unprocessable request body", nil)
		return
	}

	if err := h.s.CreateTask(&input); err != nil {
		httpresponse.FailedCreate(w)
		return
	}

	httpresponse.SendResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		httpresponse.SendErrorResponse(w, http.StatusBadRequest, "invalidId", "Invalid Id", nil)
		return
	}

	defer r.Body.Close()

	var input task.UpdateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpresponse.SendErrorResponse(w, http.StatusBadRequest, "invalidBody", "Invalid request body", nil)
		return
	}

	err = h.v.Struct(&input)
	if err != nil {
		httpresponse.SendErrorResponse(w, http.StatusUnprocessableEntity, "unprocessableBody", "Unprocessable request body", nil)
		return
	}

	if err := h.s.UpdateTask(id, &input); err != nil {
		httpresponse.FailedUpdate(w)
		return
	}

	httpresponse.SendResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		httpresponse.SendErrorResponse(w, http.StatusBadRequest, "invalidId", "Invalid Id", nil)
		return
	}

	if err := h.s.DeleteTask(id); err != nil {
		httpresponse.FailedDelete(w)
		return
	}

	httpresponse.SendResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		httpresponse.SendErrorResponse(w, http.StatusBadRequest, "invalidId", "Invalid Id", nil)
		return
	}

	task, err := h.s.GetById(id)
	if err != nil {
		httpresponse.NotFound(w)
		return
	}

	httpresponse.SendResponse(w, http.StatusOK, task)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	//TODO: поправить и переделать на пагинацию
	tasks, err := h.s.GetAll()
	if err != nil {
		http.Error(w, "tasks not found", http.StatusNotFound)
		return
	}

	httpresponse.SendResponse(w, http.StatusOK, tasks)
}

func NewHandler(s service.TaskService, v *validator.Validate) *Handler {
	return &Handler{s: s, v: v}
}
