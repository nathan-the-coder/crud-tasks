package main

import (
	"fmt"
	"net/http"

	"github.com/nathan-the-coder/crud-tasks/handlers"
	"github.com/nathan-the-coder/crud-tasks/store"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", handlers.HealthHandler)

	store := store.NewTaskStore()
	taskHandler := handlers.NewTaskHandler(&store);

	router.HandleFunc("GET /tasks", taskHandler.List)
	router.HandleFunc("POST /tasks", taskHandler.Create)
	router.HandleFunc("GET /tasks/{id}", taskHandler.Get)

	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Starting server at :8080")
	srv.ListenAndServe()
}

