package repository

import (
	"context"
	"log"
	"os"
	"review-transformer/internal/db"
	"review-transformer/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateReviewLanguage(review model.TranslateReview) error {
	collection := db.MongoClient.Database(os.Getenv("MONGO_DATABASE")).Collection("reviews")

	filter := bson.M{"id": review.ID}
	update := bson.M{
		"$set": bson.M{
			"translation": bson.M{
				"traslated_review":   review.TranslatedReview,
				"original_language":  review.OriginalLanguage,
				"translate_language": review.TranslatedLanguage,
			},
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Printf("Updated review language with ID %v", review.ID)

	return nil
}
