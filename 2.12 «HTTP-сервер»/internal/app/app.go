package app

import (
	"calendar/internal/handler"
	"calendar/internal/middleware"
	"calendar/internal/service"
	"fmt"
	"net/http"
)

func StartServer(port string) {
	mux := http.NewServeMux()
	storage := service.NewInMemoryStorage()

	mux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateEventHandler(w, r, storage)
	})
	mux.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateEventHandler(w, r, storage)
	})
	mux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteEventHandler(w, r, storage)
	})
	mux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForDayHandler(w, r, storage)
	})
	mux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForWeekHandler(w, r, storage)
	})
	mux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForMonthHandler(w, r, storage)
	})
	mux.HandleFunc("/get_events", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEventsHandler(w, r, storage)
	})

	loggedMux := middleware.LoggingMiddleware(mux)

	fmt.Printf("Starting server on port %s", port)
	http.ListenAndServe(":"+port, loggedMux)
}
