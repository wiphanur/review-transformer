package main

import (
	"context"
	"log"
	"os"
	"review-transformer/internal/db"
	"review-transformer/internal/migrations"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitMongoDB(os.Getenv("MONGO_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrations.InsertSupportLanguages(); err != nil {
		log.Fatalf("Could not insert supported languages: %v", err)
	}

	if err := migrations.InsertLanguageCodes(); err != nil {
		log.Fatalf("Could not insert language codes: %v", err)
	}

	if err := db.MongoClient.Disconnect(ctx); err != nil {
		log.Fatalf("Could not disconnect from MongoDB: %v", err)
	}
	log.Println("Disconnected from MongoDB.")
}
