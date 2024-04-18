package reviews

import (
	"encoding/json"
	"log"
	"net/http"
	"review-transformer/internal/kafkaclient"
	"review-transformer/internal/model"
	"review-transformer/internal/service"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	review := model.Review{}
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": "input need to be a valid review object"})
	}

	err = service.ReviewService(review)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Failed to send review for analysis"})
		return
	}

	translatedReview, err := service.TranslatedReviewService(review.Comments, review.ID)
	if err != nil {
		log.Printf("Could not translate review: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Review received"}`))
		return
	}

	isSupport, err := service.IsLanguageSupported(translatedReview.OriginalLanguage.Code, review.ID)
	if err != nil {
		log.Printf("Error checking supported language: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Review received"}`))
		return
	}

	if isSupport {
		byteReview, err := json.Marshal(review)
		if err != nil {
			log.Printf("Could not unmarshal review: %v", err)
		}

		topic := "sentiment-topic"
		kafkaMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          byteReview,
		}

		err = kafkaclient.Producer.Produce(kafkaMessage, nil)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		}

		kafkaclient.Producer.Flush(15 * 1000)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Review received and sent to sentiment analysis"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Review received"}`))
}
