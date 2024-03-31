package model

type TranslateReview struct {
	OriginalReview     string   `json:"original_review"`
	TranslatedReview   string   `json:"translated_review"`
	OriginalLanguage   Language `json:"original_language"`
	TranslatedLanguage Language `json:"translated_language"`
}

type Language struct {
	Code string `bson:"language_code" json:"language_code"`
	Name string `bson:"language_name" json:"language_name"`
}
