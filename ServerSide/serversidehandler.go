package ServerSideHandler

import (
	"fmt"
	"log"
	"net"
)

// either running server.go or here should create and start the server service
func createServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to", conn.RemoteAddr())
	// Example: read data from the connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Received:", string(buf[:n]))
}
