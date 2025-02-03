package ClientSideHandler

import (
	"fmt"
	"log"
	"net"
)

func createClient() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Example: send data to the server
	message := "Hello, Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Sent message:", message)
}
