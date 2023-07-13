package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"

	"golang.org/x/time/rate"
)

// Create a rate limiter with a rate of 100 requests per minute
var (
	limiter = rate.NewLimiter(rate.Limit(100), 1)
	// apiKey  = "48b111a7-e77a-44c1-9554-10b485867ac1" // Replace with your desired API key
)

func main() {
	http.HandleFunc("/api", secureMiddleware(rateLimitMiddleware(handlers.HandleAPI)))
	fmt.Println("Server listening on port 8080")

	// Read the SSL certificate and private key from GitHub Secrets
	certFile := os.Getenv("SERVER_CERT")
	keyFile := os.Getenv("SERVER_KEY")
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	// err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func secureMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enforce HTTPS redirection
		if r.TLS == nil {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusPermanentRedirect)
			return
		}

		// Check for API key in the request header or query parameter
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			apiKey = r.URL.Query().Get("api_key")
		}

		// Validate the API key
		if apiKey != os.Getenv("API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the rate limiter allows the request
		if !limiter.Allow() {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
