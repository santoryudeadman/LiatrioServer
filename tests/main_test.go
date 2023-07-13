package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"server/handlers"
	"testing"
)

func TestHandleAPI(t *testing.T) {
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the API key in the request header
	apiKey := os.Getenv("API_KEY")
	req.Header.Set("X-API-Key", apiKey)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HandleAPI)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	var response handlers.Payload
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Automate all the things!"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s' but got '%s'", expectedMessage, response.Message)
	}
}
