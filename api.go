package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJson writes a JSON response with the given status code and data.

type APIserver struct {
	listenAddr string
	store      Storage
}

func NewAPIserver(listenAddr string, store Storage) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlefunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlefunc(s.handleGetAccount))

	log.Println("JSON API SERVER RUNNING ON PORT", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)

	return nil
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case http.MethodGet:
		return s.handleGetAccount(w, r)
	case http.MethodPost:
		return s.handleCreateAccount(w, r)
	case http.MethodDelete:
		return s.handleDeleteAccount(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	return nil
}

func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	//account := NewAccount("John", "Doe")
	fmt.Println(id)

	return writeJSON(w, http.StatusOK, &Account{})
}

func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIserver) handleTransaction(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

type apifunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlefunc(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}

}
