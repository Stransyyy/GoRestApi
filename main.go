package main

import "fmt"

func main() {

	server := NewAPIserver(":8080")

	server.Run()

	fmt.Println("yeah buddy!")
}
