package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		// Read the message length prefix
		messageLengthBytes, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading message length:", err.Error())
			return
		}

		// Convert the message length to integer
		messageLength, err := strconv.Atoi(string(messageLengthBytes[:len(messageLengthBytes)-1]))
		if err != nil {
			fmt.Println("Error converting message length:", err.Error())
			return
		}

		// Read the actual message
		messageBytes := make([]byte, messageLength)
		_, err = reader.Read(messageBytes)
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			return
		}

		message := string(messageBytes)

		fmt.Println("Received message from", conn.RemoteAddr(), ":", message)

		// Send a response back to the client
		response := "Server received message: " + message

		// Send the response length prefix
		_, err = conn.Write([]byte(strconv.Itoa(len(response)) + "\n"))
		if err != nil {
			fmt.Println("Error sending response length:", err.Error())
			return
		}

		// Send the response
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error sending response:", err.Error())
			return
		}
	}
}
