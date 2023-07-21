package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

var limiter = rate.NewLimiter(1, 3) // Limit to 1 request per second with a burst capacity of 3

func main() {
	http.HandleFunc("/API", rateLimit(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s", r.RemoteAddr)
	json.NewEncoder(w).Encode(Response{
		Message:   "Automate all the things!",
		Timestamp: time.Now().Unix(),
	})
}
