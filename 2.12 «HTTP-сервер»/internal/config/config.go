package config

import "os"

func LoadConfig() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
