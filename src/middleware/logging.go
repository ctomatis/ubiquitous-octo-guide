package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lr := newLoggingResponseWriter(w)
		next.ServeHTTP(lr, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, lr)
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
	statusText string
}

func newLoggingResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK, http.StatusText(http.StatusOK)}
}

func (l *logResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.statusText = http.StatusText(code)
	l.ResponseWriter.WriteHeader(code)
}

func (l *logResponseWriter) String() string {
	return fmt.Sprintf("%d %s", l.statusCode, l.statusText)
}
