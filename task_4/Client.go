package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// Remove trailing newline character
		message = message[:len(message)-1]

		// Send the message length prefix
		_, err = conn.Write([]byte(strconv.Itoa(len(message)) + "\n"))
		if err != nil {
			fmt.Println("Error sending message length:", err.Error())
			return
		}

		// Send the actual message
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}

		// Read the response length prefix
		responseLengthBytes, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading response length:", err.Error())
			return
		}

		// Convert the response length to integer
		responseLength, err := strconv.Atoi(string(responseLengthBytes[:len(responseLengthBytes)-1]))
		if err != nil {
			fmt.Println("Error converting response length:", err.Error())
			return
		}

		// Read the response
		responseBytes := make([]byte, responseLength)
		_, err = conn.Read(responseBytes)
		if err != nil {
			fmt.Println("Error reading response:", err.Error())
			return
		}

		response := string(responseBytes)

		fmt.Println("Server response:", response)
	}
}
