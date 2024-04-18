package service

import (
	"context"
	"log"
	"review-transformer/internal/db"
	"review-transformer/internal/googleapi"
	"review-transformer/internal/model"
	"review-transformer/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
)

func ReviewService(review model.Review) error {
	review.SentimentAnalysisStatus = "pending"
	err := repository.InsertReview(review)
	if err != nil {
		log.Printf("Could not insert review: %v", err)
		return err
	}
	log.Printf("Inserted review: %v", review)

	return nil
}

func TranslatedReviewService(review string, reviewId int64) (model.TranslateReview, error) {
	translatedReview := model.TranslateReview{}
	translatedReview.ID = reviewId
	translatedReview.OriginalReview = review
	translatedReview.TranslatedLanguage = model.Language{Code: "en", Name: "English"}

	translatedText, originalLanguageCode, err := googleapi.TranslateText("en", review)
	if err != nil {
		return model.TranslateReview{}, err
	}

	translatedReview.TranslatedReview = translatedText
	translatedReview.OriginalLanguage = model.Language{Code: originalLanguageCode, Name: ""}

	collection := db.MongoClient.Database("airbnb").Collection("languages")
	filter := bson.D{{Key: "language_code", Value: originalLanguageCode}}
	err = collection.FindOne(context.Background(), filter).Decode(&translatedReview.OriginalLanguage)
	if err != nil {
		log.Printf("Could not find language with code %v: %v", originalLanguageCode, err)
		return translatedReview, err
	}

	err = repository.UpdateReviewLanguage(translatedReview)
	if err != nil {
		log.Printf("Could not update review language: %v", err)
		return translatedReview, err
	}

	return translatedReview, nil
}

func IsLanguageSupported(languageCode string, reviewId int64) (bool, error) {
	collection := db.MongoClient.Database("airbnb").Collection("supportedLanguages")
	filter := bson.D{{Key: "language_code", Value: languageCode}}
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Printf("Error while find a support laguage with code %v : %v", languageCode, err)
		return false, err
	}

	if result == nil {
		log.Printf("Language code: %v is not supported", languageCode)
		collection = db.MongoClient.Database("airbnb").Collection("reviews")
		filter = bson.D{{Key: "id", Value: reviewId}}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "sentiment_status", Value: "Not Supported Language"}}}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Printf("Could not update review sentiment status: %v", err)
			return false, err
		}
	}

	return true, nil
}
