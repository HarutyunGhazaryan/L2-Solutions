package handler

import (
	"calendar/internal/helpers"
	"calendar/internal/service"
	"net/http"
	"time"
)

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request, storage *service.InMemoryStorage) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	dateSTR := r.URL.Query().Get("date")
	if dateSTR == "" {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "missing event date"})
		return
	}

	date, err := time.Parse("2006-01-02", dateSTR)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid date format"})
		return
	}

	events := storage.GetEventsForMonth(date)
	helpers.WriteJSONResponse(w, http.StatusOK, events)

}
