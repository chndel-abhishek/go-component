package main

import (
    "encoding/json"
    "net/http"
    "github.com/rs/cors"
)

type Message struct {
    Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Hello from Go API!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(message)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api", handler)

    // Configure CORS to allow requests from any origin (for development purposes)
    handlerWithCORS := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    }).Handler(mux)

    http.ListenAndServe(":8080", handlerWithCORS)
}
