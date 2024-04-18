package demo

import (
	"encoding/json"
	"net/http"
	"review-transformer/internal/model"
	"review-transformer/internal/service"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	reviewDemo := model.ReviewDemo{}
	err := json.NewDecoder(r.Body).Decode(&reviewDemo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var translatedReview model.TranslateReview
	var sentiment model.Sentiment
	translatedReview, sentiment, err = service.DemoService(reviewDemo.Review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var demoResponse model.DemoResponse
	demoResponse.Review = translatedReview
	demoResponse.Sentiment = sentiment

	responseData, err := json.Marshal(demoResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
