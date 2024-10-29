package service

import "time"

type Event struct {
	ID    string    `json:"id,omitempty"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
