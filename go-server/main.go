package main

import (
	"log"
	"net/http"
	"os"

	"review-transformer/internal/app"
	"review-transformer/internal/db"
	"review-transformer/internal/kafkaclient"
	"review-transformer/internal/migration"

	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitMongoDB(os.Getenv("MONGO_URI"))
	migration.Migrate()
	kafkaclient.InitProducer()
	defer kafkaclient.Producer.Close()
	kafkaclient.InitConsumer()
	router := app.NewRouter()
	wrappedRouter := corsMiddleware(router)

	go kafkaclient.ReviewConsumer()

	log.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", wrappedRouter)
}
