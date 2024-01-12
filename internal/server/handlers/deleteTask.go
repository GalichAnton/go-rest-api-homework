package handlers

import (
	"net/http"

	"github.com/GalichAnton/go-rest-api-homework/internal/storage"
	"github.com/go-chi/chi/v5"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	storage.TaskStorage.M.Lock()
	defer storage.TaskStorage.M.Unlock()

	if _, ok := storage.TaskStorage.Elems[taskID]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(storage.TaskStorage.Elems, taskID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
