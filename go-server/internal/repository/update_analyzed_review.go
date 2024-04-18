package repository

import (
	"context"
	"log"
	"os"
	"review-transformer/internal/db"
	"review-transformer/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateReviewWithSentimentalScore(review model.AnalyzedReview) error {
	collection := db.MongoClient.Database(os.Getenv("MONGO_DATABASE")).Collection("reviews")

	filter := bson.M{"id": review.ID}
	update := bson.M{
		"$set": bson.M{
			"sentiment": bson.M{
				"score":              review.Sentiment.Score,
				"magnitude":          review.Sentiment.Magnitude,
				"sentiment_polarity": review.Sentiment.SentimentPolarity,
			},
			"sentiment_status": "completed",
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Printf("Updated review sentimental score with ID %v", review.ID)

	return nil
}
