package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/nathan-the-coder/crud-tasks/db"
	"github.com/nathan-the-coder/crud-tasks/handlers"
	"github.com/nathan-the-coder/crud-tasks/store"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", handlers.HealthHandler)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	sqlDB, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}
	queries := db.New(sqlDB)

	store := store.NewTaskStore(queries)

	taskHandler := handlers.NewTaskHandler(store);

	router.HandleFunc("GET /tasks", taskHandler.List)
	router.HandleFunc("POST /tasks", taskHandler.Create)
	router.HandleFunc("GET /tasks/{id}", taskHandler.Get)
	router.HandleFunc("PUT /tasks/{status}/{id}", taskHandler.Update)
	router.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	srv := http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: router,
	}

	fmt.Println("Starting server at", port)
	srv.ListenAndServe()
}

