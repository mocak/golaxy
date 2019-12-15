package controller

import (
	"encoding/json"
	"net/http"
)

// Response is the general response structure
type Response struct {
	Data interface{}
}

// successfulResponseListener sets necessary information to the response and prevents code duplication
func successfulResponseListener(w http.ResponseWriter, rsp *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(&rsp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
