package handler

import (
	"calendar/internal/helpers"
	"calendar/internal/service"
	"net/http"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request, storage *service.InMemoryStorage) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	events := storage.GetEvent()
	helpers.WriteJSONResponse(w, http.StatusOK, events)
}
