package main

import (
	"fmt"
	"net/http"
	"errors"
	"github.com/KL2412/GO-project/routes"
	"github.com/KL2412/GO-project/config"
)

func main() {
	cfg := config.New()
	router := http.NewServeMux()

	routes.RegisterRoutes(router, cfg)

	srv := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	fmt.Printf("URL Shortener is running on %s\n", cfg.BaseURL)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("An error occurred:", err)
	}
}