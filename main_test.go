package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add Basic Auth header
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("yourUsername:yourPassword")))

	rr := httptest.NewRecorder()

	// Wrap your handler function with the rate limiter
	handler := http.HandlerFunc(rateLimit(handler))

	// Call ServeHTTP method directly since we're not running a full server
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Response
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Message) == 0 {
		t.Errorf("Expected non-empty message")
	}

	if response.Timestamp < time.Now().Add(-1*time.Minute).Unix() {
		t.Errorf("Timestamp is too old")
	}
}
