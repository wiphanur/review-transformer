package repository

import (
	"context"
	"log"

	"review-transformer/internal/db"
	"review-transformer/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertReview(review model.Review) error {
	var insertResult *mongo.InsertOneResult
	var err error

	collection := db.MongoClient.Database("airbnb").Collection("reviews")

	insertResult, err = collection.InsertOne(context.TODO(), review)
	if err != nil {
		return err
	}

	log.Printf("Inserted review with ID %v", insertResult.InsertedID)
	return nil
}
