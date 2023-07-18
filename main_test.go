package main

import (
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

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
