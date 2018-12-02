package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shirobrak/newsapi/adapters"
	"github.com/shirobrak/newsapi/repositories"
	"github.com/shirobrak/newsapi/services"
)

func topicsAPIHandler(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Path[len("/topics/"):]
	repository := repositories.NewNewsAPIRepository()
	adapter := adapters.NewArticleGetter(repository)
	service := services.NewTopicsAPIService(adapter)
	serviceResponse, err := service.Run(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(serviceResponse)
}

func main() {
	// Load ENV File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Register API Handlers.
	http.HandleFunc("/topics/", topicsAPIHandler)

	// Serve
	log.Fatal(http.ListenAndServe(":8080", nil))

}
