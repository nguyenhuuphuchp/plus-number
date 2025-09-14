package main

import (
	"log"
	"net/http"
	"os"
	"plus-number/internal/handler"

	"plus-number/internal/db"
	//"plus-number/internal/model"
)

func main() {
	// init DB
	postgres, err := db.NewPostgres("postgres", "your_password", "localhost", "test", 5432)
	if err != nil {
		log.Fatalf("cannot init db: %v", err)
	}
	defer postgres.Close()

	// init handler
	addHandler := &handler.AddHandler{DB: postgres}

	mux := http.NewServeMux()
	mux.Handle("/api/add", addHandler)

	// chọn port từ ENV, mặc định 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server listening on :%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
