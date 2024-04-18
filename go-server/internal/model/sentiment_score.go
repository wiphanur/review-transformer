package model

type AnalyzedReview struct {
	ID        int64     `json:"id,omitempty" bson:"id"`
	Review    string    `json:"review"`
	Sentiment Sentiment `json:"sentiment"`
}

type Sentiment struct {
	Score             float64 `json:"score" bson:"score"`
	Magnitude         float64 `json:"magnitude" bson:"magnitude"`
	SentimentPolarity string  `json:"sentimentPolarity" bson:"sentiment_polarity"`
}
