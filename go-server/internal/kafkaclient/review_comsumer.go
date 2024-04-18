package kafkaclient

import (
	"fmt"
	"log"
)

func ReviewConsumer() {
	Consumer.SubscribeTopics([]string{"review-topic", "analyzed-sentiment-reviews-topic"}, nil)

	for {
		msg, err := Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		switch *msg.TopicPartition.Topic {
		case "analyzed-sentiment-reviews-topic":
			analyedReviewConsumer(msg)
		case "review-topic":
			processReviewMessage(msg)
		default:
			fmt.Printf("Unknown topic: %s\n", *msg.TopicPartition.Topic)
		}
	}
}
