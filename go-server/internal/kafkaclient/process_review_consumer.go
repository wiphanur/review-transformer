package kafkaclient

import (
	"encoding/json"
	"log"
	"review-transformer/internal/model"
	"review-transformer/internal/service"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func processReviewMessage(msg *kafka.Message) error {
	var reviewMessage model.Review
	if err := json.Unmarshal(msg.Value, &reviewMessage); err != nil {
		log.Printf("Could not unmarshal review: %v", err)
		return err
	}

	err := service.ReviewService(reviewMessage)
	if err != nil {
		log.Printf("Could not process review: %v", err)
		return err
	}

	translatedReview, err := service.TranslatedReviewService(reviewMessage.Comments, reviewMessage.ID)
	if err != nil {
		log.Printf("Could not translate review: %v", err)
		return err
	}

	isSupport, err := service.IsLanguageSupported(translatedReview.OriginalLanguage.Code, reviewMessage.ID)
	if err != nil {
		log.Printf("Error checking supported language: %v", err)
		return err
	}

	if isSupport {
		byteReview, err := json.Marshal(reviewMessage)
		if err != nil {
			log.Printf("Could not unmarshal review: %v", err)
		}

		topic := "sentiment-topic"
		kafkaMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          byteReview,
		}

		err = Producer.Produce(kafkaMessage, nil)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		}

		Producer.Flush(15 * 1000)
	}

	return nil
}
