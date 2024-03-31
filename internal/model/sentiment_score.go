package model

type AnalyzedReview struct {
	Review    string    `json:"review"`
	Sentiment Sentiment `json:"sentiment"`
}

type Sentiment struct {
	Score             float64 `json:"score"`
	Magnitude         float64 `json:"magnitude"`
	SentimentPolarity string  `json:"sentimentPolarity"`
}
