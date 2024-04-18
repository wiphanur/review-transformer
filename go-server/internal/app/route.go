package app

import (
	"fmt"
	"net/http"
	"review-transformer/internal/app/demo"
	"review-transformer/internal/app/reviews"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/health", healthCheckHandler)
	router.HandleFunc("/demo", demo.DemoHandler)
	router.HandleFunc("/review", reviews.ReviewHandler)
	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is up and running!")
}
