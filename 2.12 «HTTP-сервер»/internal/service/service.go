package service

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type InMemoryStorage struct {
	mu     sync.Mutex
	events map[string]Event
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		events: make(map[string]Event),
	}
}

func (ms *InMemoryStorage) CreateEvent(event Event) string {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	event.ID = uuid.New().String()
	ms.events[event.ID] = event
	return event.Title
}

func (ms *InMemoryStorage) UpdateEvent(updatedEvent Event) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.events[updatedEvent.ID]; exists {
		ms.events[updatedEvent.ID] = updatedEvent
		return true
	}
	return false
}

func (ms *InMemoryStorage) DeleteEvent(id string) bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	_, exists := ms.events[id]
	if exists {
		delete(ms.events, id)
		return true
	}
	return false
}

func (ms *InMemoryStorage) GetEvent() []Event {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	allEvents := make([]Event, 0, len(ms.events))
	for _, event := range ms.events {
		allEvents = append(allEvents, event)
	}
	return allEvents
}

func (ms *InMemoryStorage) GetEventsForDay(date time.Time) []Event {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var dailyEvents []Event
	for _, event := range ms.events {
		if event.Date.Equal(date) {
			dailyEvents = append(dailyEvents, event)
		}
	}
	return dailyEvents
}

func (ms *InMemoryStorage) GetEventsForWeek(date time.Time) []Event {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	startOfWeek := date.AddDate(0, 0, -int(date.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var weeklyEvents []Event
	for _, event := range ms.events {
		if event.Date.Before(endOfWeek) && event.Date.After(startOfWeek) {
			weeklyEvents = append(weeklyEvents, event)
		}
	}
	return weeklyEvents
}

func (ms *InMemoryStorage) GetEventsForMonth(date time.Time) []Event {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var monthlyEvents []Event
	for _, event := range ms.events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			monthlyEvents = append(monthlyEvents, event)
		}
	}
	return monthlyEvents
}
