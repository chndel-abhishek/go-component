package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Message struct {
	Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello from Go API!"}
	w.Header().Set("Content-Type", "application/json")

	// Check error from json.Encode and handle it
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handler)

	// Configure CORS to allow requests from any origin (for development purposes)
	handlerWithCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	// Check error from http.ListenAndServe and handle it
	err := http.ListenAndServe(":8080", handlerWithCORS)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}