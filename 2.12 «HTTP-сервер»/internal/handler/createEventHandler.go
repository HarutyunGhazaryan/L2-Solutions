package handler

import (
	"calendar/internal/helpers"
	"calendar/internal/service"
	"net/http"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, storage *service.InMemoryStorage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	params, err := helpers.ParseAndValidateEvent(r)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	event := service.Event{
		Title: params["title"].(string),
		Date:  params["date"].(time.Time),
	}

	createdEvent := storage.CreateEvent(event)
	helpers.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"result": "event created", "event": createdEvent})
}
