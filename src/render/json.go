package render

import (
	"encoding/json"
	"net/http"

	"go.vemo/src/settings"
)

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   any    `json:"data,omitempty"`
	Errors any    `json:"error,omitempty"`
}

func rJSON(w http.ResponseWriter, r Response) {
	r.Status = http.StatusText(r.Code)
	w.Header().Set("content-type", settings.Get("content_type"))
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

func Send(w http.ResponseWriter, body any, statusCode int) {
	res := Response{
		Code: statusCode,
		Data: body,
	}
	rJSON(w, res)
}

func Abort(w http.ResponseWriter, body any, statusCode int) {
	res := Response{
		Code:   statusCode,
		Errors: body,
	}
	rJSON(w, res)
}
