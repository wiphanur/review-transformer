package model

type DemoResponse struct {
	Review    TranslateReview `json:"review"`
	Sentiment Sentiment       `json:"sentiment"`
}
