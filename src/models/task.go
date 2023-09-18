package models

import (
	"errors"
	"strings"

	"go.vemo/src/render"
)

type Task struct {
	Id     int             `json:"id"`
	Name   string          `json:"name"`
	Detail string          `json:"detail"`
	Done   bool            `json:"done"`
	Due    render.DateTime `json:"due"`
}

type Tasks []Task

func (t *Task) Check() (fields map[string]string, err error) {
	fields = map[string]string{}
	if len(t.Name) == 0 {
		fields["name"] = "required value"
	}
	if len(t.Detail) == 0 {
		fields["detail"] = "required value"
	}

	if strings.Contains(t.Due.String(), "0001-01-01") {
		fields["due"] = "required value"
	}

	if len(fields) > 0 {
		return fields, errors.New("fields with error")
	}
	return
}
