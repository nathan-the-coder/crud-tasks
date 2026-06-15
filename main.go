package main

import (
	"fmt"
	"net/http"

	"github.com/nathan-the-coder/crud-tasks/handlers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/health", handlers.HealthHandler)

	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Starting server at :8080")
	srv.ListenAndServe()
}

