package handler

import (
	"calendar/internal/helpers"
	"calendar/internal/service"
	"net/http"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, storage *service.InMemoryStorage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "missing event ID"})
		return
	}
	found := storage.DeleteEvent(id)
	if found {
		helpers.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"result": "event deleted"})
	} else {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "id not found"})
	}
}
