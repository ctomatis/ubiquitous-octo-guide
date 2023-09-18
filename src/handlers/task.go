package handlers

import (
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"go.vemo/src/models"
	"go.vemo/src/render"
	"go.vemo/src/repo"
	"go.vemo/src/validator"
)

var (
	itemPathRgx = regexp.MustCompile(`^\/tasks\/(\d+)$`)
	listPathRgx = regexp.MustCompile(`^\/tasks[\/]*$`)
)

type TaskHandler struct {
	*repo.Tasks
}

func ParamId(s string, id *int) bool {
	s = strings.Trim(s, "/")
	parts := strings.Split(s, "/")
	if len(parts) == 2 {
		if val, err := strconv.Atoi(parts[1]); err == nil && val > 0 {
			*id = val
			return true
		}
	}
	return false
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && itemPathRgx.MatchString(r.URL.Path):
		h.Find(w, r)
		return
	case r.Method == http.MethodGet && listPathRgx.MatchString(r.URL.Path):
		h.All(w, r)
		return
	case r.Method == http.MethodPost && listPathRgx.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	case r.Method == http.MethodPut && itemPathRgx.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	case r.Method == http.MethodDelete && itemPathRgx.MatchString(r.URL.Path):
		h.Delete(w, r)
		return
	default:
		render.Abort(w, nil, http.StatusNotFound)
		return
	}
}

func (h *TaskHandler) Find(w http.ResponseWriter, r *http.Request) {
	var id int
	if !ParamId(r.URL.Path, &id) {
		render.Abort(w, nil, http.StatusBadRequest)
		return
	}

	task, err := h.Tasks.Find(id)
	if err != nil {
		render.Abort(w, nil, http.StatusNotFound)
		return
	}
	render.Send(w, task, http.StatusOK)
}

func (h *TaskHandler) All(w http.ResponseWriter, r *http.Request) {
	render.Send(w, h.Tasks.All(), http.StatusOK)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if m, err := validator.Validate(r.Body, &task); err != nil {
		if err == io.EOF || len(m) == 0 {
			render.Abort(w, nil, http.StatusBadRequest)
			return
		}
		render.Abort(w, m, http.StatusUnprocessableEntity)
		return
	}
	h.Tasks.Create(&task)
	render.Send(w, task, http.StatusCreated)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var id int
	if !ParamId(r.URL.Path, &id) {
		render.Abort(w, nil, http.StatusBadRequest)
		return
	}
	_, err := h.Tasks.Find(id)
	if err != nil {
		render.Abort(w, nil, http.StatusNotFound)
		return
	}

	var task models.Task
	if m, err := validator.Validate(r.Body, &task); err != nil {
		if err == io.EOF || len(m) == 0 {
			render.Abort(w, nil, http.StatusBadRequest)
			return
		}
		render.Abort(w, m, http.StatusUnprocessableEntity)
		return
	}

	h.Tasks.Update(id, &task)
	render.Send(w, task, http.StatusOK)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var id int
	if !ParamId(r.URL.Path, &id) {
		render.Abort(w, nil, http.StatusBadRequest)
		return
	}
	_, err := h.Tasks.Find(id)
	if err != nil {
		render.Abort(w, nil, http.StatusNotFound)
		return
	}
	h.Tasks.Delete(id)
	render.Send(w, nil, http.StatusNoContent)
}
