package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GalichAnton/go-rest-api-homework/internal/storage"
)

func TasksHandler(w http.ResponseWriter, _ *http.Request) {
	resp, err := json.Marshal(storage.TaskStorage.Elems)
	if err != nil {
		http.Error(w, "Failed to decode tasks data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)
}
