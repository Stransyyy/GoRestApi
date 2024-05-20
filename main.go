package main

import (
	"log"
)

func main() {
	store, err := NewPosrgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIserver(":3000", store)
	server.Run()

}
