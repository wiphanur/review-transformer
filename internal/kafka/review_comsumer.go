package kafka

import (
	"encoding/json"
	"log"
	"os"

	"review-transformer/internal/model"
	"review-transformer/internal/repository"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConsumeReviews() {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          "review-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	defer kafkaConsumer.Close()

	kafkaConsumer.SubscribeTopics([]string{"review-topic"}, nil)

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue // Continue to the next message
		}

		var reviewMessage model.Review
		if err := json.Unmarshal(msg.Value, &reviewMessage); err != nil {
			log.Printf("Could not unmarshal review: %v", err)
			continue // Continue to the next message
		}

		// Insert the review into the MongoDB database
		reviewMessage.SentimentAnalysisStatus = "pending"
		err = repository.InsertReview(reviewMessage)
		if err != nil {
			log.Printf("Could not insert review: %v", err)
			continue // Continue to the next message
		}
		log.Printf("Inserted review: %v", reviewMessage)

		// Send data to sentiment analysis service through a Kafka topic
		var kafkaProducer *kafka.Producer
		kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS")})
		if err != nil {
			log.Fatalf("Failed to create producer: %s", err)
		}

		if err := json.Unmarshal(msg.Value, &reviewMessage); err != nil {
			log.Printf("Could not unmarshal review: %v", err)
			continue // Continue to the next message
		}

		topic := "sentiment-topic"
		kafkaProducer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          msg.Value,
		}, nil)

		kafkaProducer.Flush(15 * 1000)
		kafkaProducer.Close()

	}
}
