package main

import (
	"encoding/json"
	"expvar"
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"golang.org/x/time/rate"
)

var (
	limiter = rate.NewLimiter(1, 3)     // Limit to 1 request per second with a burst capacity of 3
	counter = expvar.NewInt("requests") // Counter for total number of requests
)

type Response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	expvar.Publish("time", expvar.Func(func() interface{} { return time.Now().Unix() }))

	// Create a serve mux just for the profiling and expvar endpoints
	mux := http.NewServeMux()
	// Wrap the pprof handlers with the auth function
	mux.HandleFunc("/debug/pprof/", auth(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", auth(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", auth(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", auth(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", auth(pprof.Trace))
	// Wrap the expvar handler with the auth function
	mux.HandleFunc("/debug/vars", auth(expvar.Handler().ServeHTTP))

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", mux))
	}()

	http.HandleFunc("/", rateLimit(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "yourUsername" || password != "yourPassword" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	counter.Add(1) // Increment counter
	log.Printf("Received request from %s", r.RemoteAddr)
	json.NewEncoder(w).Encode(Response{
		Message:   "Automate all the things!",
		Timestamp: time.Now().Unix(),
	})
}
