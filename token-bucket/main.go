package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	message := Message{
		Status: "Successful",
		Body:   "Hi, you have reached the API. How may I help you?",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

func main() {
	http.Handle("/ping", rateLimiter(http.HandlerFunc(endpointHandler)))

	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port 8080:", err)
	}
}
