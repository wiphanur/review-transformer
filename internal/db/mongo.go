package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Client is the MongoDB client that will be used throughout your application.
	MongoClient *mongo.Client
)

// InitMongoDB initializes the MongoDB client.
func InitMongoDB(connectionString string) {
	var err error
	clientOptions := options.Client().ApplyURI(connectionString)
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}
