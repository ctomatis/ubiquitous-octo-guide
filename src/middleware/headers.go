package middleware

import (
	"net/http"

	"go.vemo/src/render"
	"go.vemo/src/settings"
)

func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !settings.IsDebug() {
			if contentType := r.Header.Get("Content-Type"); contentType != settings.Get("content_type") {
				render.Abort(w, nil, http.StatusUnsupportedMediaType)
				return
			}

			if auth := r.Header.Get("Authorization"); auth != settings.Get("api_key") {
				render.Abort(w, nil, http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
