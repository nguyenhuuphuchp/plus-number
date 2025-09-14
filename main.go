//package main
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"os"
//	"time"
//
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//type requestBody struct {
//	A int `json:"a"`
//	B int `json:"b"`
//}
//
//type responseBody struct {
//	Result int `json:"result"`
//}
//
//var db *pgxpool.Pool
//
//func main() {
//	// Lấy DATABASE_URL từ biến môi trường
//	databaseUrl := os.Getenv("DATABASE_URL")
//	if databaseUrl == "" {
//		log.Fatal("DATABASE_URL environment variable is required, ví dụ: postgres://user:pass@localhost:5432/dbname")
//	}
//
//	// Kết nối Postgres
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	pool, err := pgxpool.New(ctx, databaseUrl)
//	if err != nil {
//		log.Fatalf("failed to create db pool: %v", err)
//	}
//	db = pool
//	defer db.Close()
//
//	// API endpoint
//	http.HandleFunc("/api/add", addHandler)
//
//	port := os.Getenv("PORT")
//	if port == "" {
//		port = "8080"
//	}
//	log.Printf("Server listening on :%s ...", port)
//	log.Fatal(http.ListenAndServe(":"+port, nil))
//}
//
//func addHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	var req requestBody
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
//		return
//	}
//
//	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
//	defer cancel()
//
//	var sum int
//	// Gọi function add_numbers_hard trong PostgreSQL
//	err := db.QueryRow(ctx, "SELECT add_numbers_hard($1, $2)", req.A, req.B).Scan(&sum)
//	if err != nil {
//		http.Error(w, fmt.Sprintf("DB error: %v", err), http.StatusInternalServerError)
//		return
//	}
//
//	resp := responseBody{Result: sum}
//	w.Header().Set("Content-Type", "application/json")
//	_ = json.NewEncoder(w).Encode(resp)
//}
