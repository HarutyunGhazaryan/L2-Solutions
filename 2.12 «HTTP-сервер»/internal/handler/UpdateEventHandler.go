package handler

import (
	"calendar/internal/helpers"
	"calendar/internal/service"
	"net/http"
	"time"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request, storage *service.InMemoryStorage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	params, err := helpers.ParseAndValidateUpdateEventParams(r)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	event := service.Event{
		ID:    params["id"].(string),
		Title: params["title"].(string),
		Date:  params["date"].(time.Time),
	}

	found := storage.UpdateEvent(event)
	if found {
		helpers.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"result": "event updated"})
	} else {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "id not found"})
	}
}
