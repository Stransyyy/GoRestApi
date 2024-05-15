package main

import "net/http"

type APIserver struct {
	listenAddr string
}

func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIserver) HandleAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIserver) HandleGetAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIserver) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIserver) HandleDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	return nil
}
func (s *APIserver) HandleTransaction(w http.ResponseWriter, r *http.Request) error {

	return nil
}
