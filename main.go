package main

func main() {

	server := NewAPIserver(":8080")
	server.Run()

}
