package config

import (
	"os"
)

type Config struct {
	BaseURL string
	Port    string
}

func New() *Config {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		BaseURL: baseURL,
		Port:    port,
	}
} 