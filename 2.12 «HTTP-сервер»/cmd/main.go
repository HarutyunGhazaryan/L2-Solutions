package main

import (
	"calendar/internal/app"
	"calendar/internal/config"
)

func main() {
	port := config.LoadConfig()
	app.StartServer(port)
}
