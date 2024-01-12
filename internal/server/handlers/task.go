package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GalichAnton/go-rest-api-homework/internal/storage"
	"github.com/go-chi/chi/v5"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	storage.TaskStorage.M.RLock()
	defer storage.TaskStorage.M.RUnlock()

	task, ok := storage.TaskStorage.Elems[taskID]
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Failed to encode task data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}
