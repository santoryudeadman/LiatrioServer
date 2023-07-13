package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Payload struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	payload := Payload{
		Message:   "Automate all the things!",
		Timestamp: time.Now().Unix(),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPayload)
}
