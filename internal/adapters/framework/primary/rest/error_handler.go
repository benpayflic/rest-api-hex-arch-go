package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func (r *ErrorResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Return a http error response
func httpErrorResponse(code int, message string, w http.ResponseWriter, err error) {
	var errResp ErrorResponse
	errResp.StatusCode = code
	errResp.Message = message
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(errResp)
	log.Println(err)
}
