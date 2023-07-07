package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
}

var users = make(map[string]User)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Change Password")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			registerUser(reader)
		case "2":
			loginUser(reader)
		case "3":
			changePassword(reader)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func registerUser(reader *bufio.Reader) {
	fmt.Println("=== Register ===")

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	username = strings.TrimSpace(username)

	if _, exists := users[username]; exists {
		fmt.Println("Username already exists. Please try again.")
		return
	}

	fmt.Print("Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	password = strings.TrimSpace(password)

	users[username] = User{Username: username, Password: password}
	fmt.Println("Registration successful!")
}

func loginUser(reader *bufio.Reader) {
	fmt.Println("=== Login ===")

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	password = strings.TrimSpace(password)

	user, exists := users[username]
	if !exists || user.Password != password {
		fmt.Println("Invalid username or password.")
		return
	}

	fmt.Println("Login successful!")
}

func changePassword(reader *bufio.Reader) {
	fmt.Println("=== Change Password ===")

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	username = strings.TrimSpace(username)

	user, exists := users[username]
	if !exists {
		fmt.Println("User does not exist.")
		return
	}

	fmt.Print("Enter current password: ")
	currentPassword, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	currentPassword = strings.TrimSpace(currentPassword)

	if user.Password != currentPassword {
		fmt.Println("Invalid current password.")
		return
	}

	fmt.Print("Enter new password: ")
	newPassword, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	newPassword = strings.TrimSpace(newPassword)

	user.Password = newPassword
	users[username] = user

	fmt.Println("Password changed successfully!")
}
