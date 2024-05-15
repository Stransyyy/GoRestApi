package api

import (
	"encoding/json"
	"net/http"
)

// TaskHandler is the handler for the /task endpoint
type TaskParams struct {
	Username string
}

// TaskHandler is the handler for the /task endpoint reponse
type TaskResponse struct {
	// success code
	Code int `json:"code"`
	// task message
	Result string `json:"result"`
}

type Error struct {
	Code int

	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {

	reps := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(reps)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
