package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the HomePage!")
	}).Methods(http.MethodGet)

	log.Println("listening...")

	log.Fatal(http.ListenAndServe(":8081", router))
}
