package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

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
	message := Message{
		Status: "Request Failed",
		Body:   "The API is at capacity, try again later.",
	}
	jsonMessage, _ := json.Marshal(message)

	// Define the rate limiter (1 request per second)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))

	// Apply the rate limiter to the endpoint
	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))

	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port 8080:", err)
	}
}
