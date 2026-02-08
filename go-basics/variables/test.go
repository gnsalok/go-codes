package main

import (
	"fmt"
)

// User struct to represent a user
type User struct {
	ID   int
	Name string
}

// Users struct to hold a slice of users and a method to check emptiness
type Users struct {
	users []User
}

func main() {
	// Create an empty Users struct
	var users Users

	// Add a user to the slice
	// user := User{ID: 1, Name: "John Doe"}
	// users.users = append(users.users, user)
	users.users = nil
	if len(users.users) > 0 {
		fmt.Println("Users found:", users.users)
	} else {
		print("slice is null")
	}
}
