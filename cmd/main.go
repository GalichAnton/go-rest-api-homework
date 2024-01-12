package main

import (
	"fmt"
	"net/http"

	"github.com/GalichAnton/go-rest-api-homework/internal/server/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.TasksHandler)
	r.Post("/tasks", handlers.PostTaskHandler)
	r.Get("/tasks/{id}", handlers.TaskHandler)
	r.Delete("/tasks/{id}", handlers.DeleteTaskHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
