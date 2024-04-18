package kafkaclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var Consumer *kafka.Consumer

func InitConsumer() {
	createTopics()

	var err error
	Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          "review-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
}

func createTopics() {

	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	topic := []string{"review-topic", "sentiment-topic", "analyzed-sentiment-reviews-topic"}
	// Create a new AdminClient.
	// AdminClient can also be instantiated using an existing
	// Producer or Consumer instance, see NewAdminClientFromProducer and
	// NewAdminClientFromConsumer.
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	defer a.Close()
	for _, element := range topic {
		// Contexts are used to abort or limit the amount of time
		// the Admin call blocks waiting for a result.
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Create topics on cluster.
		// Set Admin options to wait for the operation to finish (or at most 60s)
		maxDur, err := time.ParseDuration("60s")
		if err != nil {
			panic("ParseDuration(60s)")
		}
		results, err := a.CreateTopics(
			ctx,
			// Multiple topics can be created simultaneously
			// by providing more TopicSpecification structs here.
			[]kafka.TopicSpecification{{
				Topic:             element,
				NumPartitions:     1,
				ReplicationFactor: 1}},
			// Admin options
			kafka.SetAdminOperationTimeout(maxDur))
		if err != nil {
			fmt.Printf("Failed to create topic: %v\n", err)
			os.Exit(1)
		}

		// Print results
		for _, result := range results {
			fmt.Printf("%s\n", result)
		}
	}
}
