package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func WriteJSONResponse(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func ParseAndValidateEvent(r *http.Request) (map[string]interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.New("invalid form date")
	}

	title := r.FormValue("title")
	if title == "" {
		return nil, errors.New("title is required")
	}

	dateStr := r.FormValue("date")
	if dateStr == "" {
		return nil, errors.New("date is required")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-FF")
	}

	params := map[string]interface{}{
		"title": title,
		"date":  date,
	}

	return params, nil
}

func ParseAndValidateUpdateEventParams(r *http.Request) (map[string]interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, errors.New("invalid form date")
	}

	eventID := r.FormValue("id")
	if eventID == "" {
		return nil, errors.New("id is required")
	}

	params, err := ParseAndValidateEvent(r)
	if err != nil {
		return nil, err
	}

	params["id"] = eventID
	return params, nil
}
