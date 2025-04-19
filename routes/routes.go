package routes

import (
	"net/http"
	"github.com/KL2412/GO-project/urlshortener"
	"github.com/KL2412/GO-project/config"
)

func RegisterRoutes(router *http.ServeMux, cfg *config.Config) {
	shortener := urlshortener.New(cfg)
	handlers := urlshortener.NewHandlers(shortener, cfg)
	
	router.HandleFunc("GET /{$}", handlers.HandleIndex)
	router.HandleFunc("POST /shorten", handlers.HandleShortenWithTemplate)
	router.HandleFunc("GET /short/", shortener.HandleRedirect)
} 