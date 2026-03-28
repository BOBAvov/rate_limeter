package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Health struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	health := Health{
		Status: "ok",
		Time:   time.Now(),
	}
	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
