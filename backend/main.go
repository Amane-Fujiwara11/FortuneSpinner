package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure/mysql"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure/repository"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/interface/handler"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/gacha"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/point"
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

	// Initialize database connection
	db, err := mysql.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	gachaRepo := repository.NewGachaRepository(db)
	pointRepo := repository.NewPointRepository(db)

	// Initialize usecases
	gachaUsecase := gacha.NewGachaUsecase(gachaRepo, pointRepo, userRepo)
	pointUsecase := point.NewPointUsecase(pointRepo, userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userRepo)
	gachaHandler := handler.NewGachaHandler(gachaUsecase)
	pointHandler := handler.NewPointHandler(pointUsecase)

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
	mux.HandleFunc("/api/users", corsHandler(userHandler.CreateUser))

	// Gacha routes
	mux.HandleFunc("/api/gacha/execute", corsHandler(gachaHandler.ExecuteGacha))
	mux.HandleFunc("/api/gacha/history", corsHandler(gachaHandler.GetGachaHistory))

	// Point routes
	mux.HandleFunc("/api/points/balance", corsHandler(pointHandler.GetBalance))
	mux.HandleFunc("/api/points/transactions", corsHandler(pointHandler.GetTransactionHistory))

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