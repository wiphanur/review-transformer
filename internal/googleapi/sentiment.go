package googleapi

import (
	"context"
	"log"
	"review-transformer/internal/model"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
)

func SentimantalAnalysis(review string) (model.Sentiment, error) {
	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		return model.Sentiment{}, err
	}
	defer client.Close()

	// var sentiment model.Sentiment
	sentimentResponse, err := analyzeSentiment(ctx, client, review)
	if err != nil {
		return model.Sentiment{}, err
	}

	var sentiment model.Sentiment
	sentiment.Magnitude = float64(sentimentResponse.DocumentSentiment.Magnitude)
	sentiment.Score = float64(sentimentResponse.DocumentSentiment.Score)

	log.Printf("response from Google Language API: %v", sentiment)
	return sentiment, nil
}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})

	if err != nil {
		return nil, err
	}

	return sentiment, nil
}
