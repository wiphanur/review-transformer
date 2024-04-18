package kafkaclient

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var Producer *kafka.Producer

func InitProducer() {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS")})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
}
