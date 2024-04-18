package migration

import (
	"log"
	"os"
	"review-transformer/internal/db"
	"review-transformer/internal/migrations"
)

func Migrate() {
	db.InitMongoDB(os.Getenv("MONGO_URI"))

	log.Println("Start migration..")
	if err := migrations.InsertSupportLanguages(); err != nil {
		log.Fatalf("Could not insert supported languages: %v", err)
	}

	if err := migrations.InsertLanguageCodes(); err != nil {
		log.Fatalf("Could not insert language codes: %v", err)
	}

	log.Println("Finish migration..")
}
