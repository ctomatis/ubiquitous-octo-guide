package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.vemo/src/handlers"
	"go.vemo/src/repo"
)

func handler() *handlers.TaskHandler {
	repo := &repo.Tasks{}
	repo.Init()

	return &handlers.TaskHandler{Tasks: repo}
}
func TestTaskList(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	res := httptest.NewRecorder()

	handler().All(res, req)

	if res.Result().StatusCode != http.StatusOK {
		t.Error(res.Body.String())
	}
}

func TestTaskGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	res := httptest.NewRecorder()

	handler().Find(res, req)

	if res.Result().StatusCode != http.StatusOK {
		t.Error(res.Body.String())
	}
}

func TestTaskPost(t *testing.T) {

	payload := strings.NewReader(`{
		"name": "A new task",
		"detail": "Create a new task",
		"done": false,
		"due": "2023-11-10 10:00:00"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/tasks", payload)
	res := httptest.NewRecorder()

	handler().Create(res, req)

	if res.Result().StatusCode != http.StatusCreated {
		t.Error(res.Body.String())
	}
}

func TestTaskPut(t *testing.T) {

	payload := strings.NewReader(`{
		"name": "Update a task",
		"detail": "Detail update",
		"done": false,
		"due": "2023-11-01 10:00:00"
	}`)

	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", payload)
	res := httptest.NewRecorder()

	handler().Update(res, req)

	if res.Result().StatusCode != http.StatusOK {
		t.Error(res.Body.String())
	}
}

func TestTasDelete(t *testing.T) {
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	res := httptest.NewRecorder()

	handler().Delete(res, req)

	if res.Result().StatusCode != http.StatusNoContent {
		t.Error(res.Body.String())
	}
}
