package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GalichAnton/go-rest-api-homework/internal/model"
	"github.com/GalichAnton/go-rest-api-homework/internal/storage"
)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := &model.Task{}
	if err := json.NewDecoder(r.Body).Decode(task); err != nil {
		http.Error(w, "Failed to decode task data", http.StatusBadRequest)
		return
	}

	storage.TaskStorage.M.Lock()
	defer storage.TaskStorage.M.Unlock()

	storage.Tasks[task.ID] = *task

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
