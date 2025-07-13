package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure/mysql"
)

func main() {
	// Database configuration
	dbConfig := mysql.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "rootpassword"),
		Database: getEnv("DB_NAME", "fortunespinner"),
	}

	// Initialize DI container with all dependencies
	container, err := infrastructure.NewContainer(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	defer container.Close()

	// Setup routes
	mux := http.NewServeMux()

	// CORS middleware
	corsHandler := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next(w, r)
		}
	}

	// User routes
	mux.HandleFunc("/api/users", corsHandler(container.UserHandler.CreateUser))

	// Gacha routes
	mux.HandleFunc("/api/gacha/execute", corsHandler(container.GachaHandler.ExecuteGacha))
	mux.HandleFunc("/api/gacha/history", corsHandler(container.GachaHandler.GetGachaHistory))

	// Point routes
	mux.HandleFunc("/api/points/balance", corsHandler(container.PointHandler.GetBalance))
	mux.HandleFunc("/api/points/transactions", corsHandler(container.PointHandler.GetTransactionHistory))

	// Health check
	mux.HandleFunc("/health", corsHandler(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}