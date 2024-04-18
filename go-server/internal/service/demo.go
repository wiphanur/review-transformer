package service

import (
	"context"
	"log"
	"review-transformer/internal/db"
	"review-transformer/internal/googleapi"
	"review-transformer/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DemoService(review string) (model.TranslateReview, model.Sentiment, error) {
	translatedReview := model.TranslateReview{}
	translatedReview.OriginalReview = review
	translatedReview.TranslatedLanguage = model.Language{Code: "en", Name: "English"}

	translatedText, originalLanguageCode, err := googleapi.TranslateText("en", review)
	if err != nil {
		return model.TranslateReview{}, model.Sentiment{}, err
	}

	translatedReview.TranslatedReview = translatedText
	translatedReview.OriginalLanguage = model.Language{Code: originalLanguageCode, Name: ""}

	collection := db.MongoClient.Database("airbnb").Collection("languages")
	filter := bson.D{{Key: "language_code", Value: originalLanguageCode}}
	err = collection.FindOne(context.Background(), filter).Decode(&translatedReview.OriginalLanguage)
	if err != nil {
		log.Printf("Could not find language with code %v: %v", originalLanguageCode, err)
		return translatedReview, model.Sentiment{}, err
	}

	collection = db.MongoClient.Database("airbnb").Collection("supportedLanguages")
	filter = bson.D{{Key: "language_code", Value: originalLanguageCode}}
	err = collection.FindOne(context.Background(), filter).Decode(&translatedReview.OriginalLanguage)
	if err != nil {
		log.Printf("This language with code %v is not supported: %v", originalLanguageCode, err)
		return translatedReview, model.Sentiment{}, err
	}

	var sentiment model.Sentiment
	sentiment, err = googleapi.SentimantalAnalysis(translatedReview.OriginalReview)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return translatedReview, model.Sentiment{}, nil
		}

		return translatedReview, model.Sentiment{}, err

	}

	if sentiment != (model.Sentiment{}) {
		sentiment.SentimentPolarity = "Neutral"
		if sentiment.Score < -0.25 {
			sentiment.SentimentPolarity = "Negative"
		}
		if sentiment.Score > 0.25 {
			sentiment.SentimentPolarity = "Positive"
		}
	}

	return translatedReview, sentiment, nil
}
