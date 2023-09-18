package main

import (
	"log"
	"net/http"
	"os"

	"go.vemo/src/handlers"
	m "go.vemo/src/middleware"
	"go.vemo/src/repo"
)

func main() {
	mux := http.NewServeMux()

	repo := &repo.Tasks{}
	repo.Init()

	task_handler := m.Logging(m.Headers(&handlers.TaskHandler{Tasks: repo}))

	mux.Handle("/tasks/", task_handler)

	run(mux)
}

func run(mux *http.ServeMux) {
	port := port()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Listening HTTP on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func port() (port string) {
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}
	return
}
