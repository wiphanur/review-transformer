package kafkaclient

import (
	"encoding/json"
	"log"

	"review-transformer/internal/model"
	"review-transformer/internal/repository"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func analyedReviewConsumer(msg *kafka.Message) error {

	var review model.AnalyzedReview
	if err := json.Unmarshal(msg.Value, &review); err != nil {
		log.Printf("Could not unmarshal review: %v", err)
		return err
	}

	err := repository.UpdateReviewWithSentimentalScore(review)
	if err != nil {
		log.Fatalf("Failed to update review: %s", err)
		return err
	}

	return nil
}
