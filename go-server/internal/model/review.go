package model

type Review struct {
	ListingID               int64  `json:"listingId" bson:"listing_id"`
	ID                      int64  `json:"id" bson:"id"`
	Date                    string `json:"date" bson:"date"`
	ReviewerID              int64  `json:"reviewerId" bson:"reviewer_id"`
	ReviewerName            string `json:"reviewerName" bson:"reviewer_name"`
	Comments                string `json:"comments" bson:"comments"`
	SentimentAnalysisStatus string `json:"sentimentStatus,omitempty" bson:"sentiment_status"`
}
