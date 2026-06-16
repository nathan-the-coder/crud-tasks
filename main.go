package main

import (
		_ "github.com/go-sql-driver/mysql"

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

	sqlDB, err := sql.Open("mysql", "root:lein2324@/task_manager?parseTime=true")
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

	srv := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Println("Starting server at :8080")
	srv.ListenAndServe()
}

