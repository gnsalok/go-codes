package main

import (
	"fmt"
	"log"
	"reflect"
)

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

// Mock objects meet the interface requirements of,
// and stand in for, more complex real ones
type registrationPreChecker interface {
	userExists(string) bool
}

type regPreCheck struct{}

func (r regPreCheck) userExists(email string) bool {
	return UserExists(email)
}

var regPreCond registrationPreChecker

func init() {
	regPreCond = regPreCheck{}

}

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User) error {
	// check if user is already registered
	found := regPreCond.userExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}

func main() {
	fmt.Println("Type of ", reflect.TypeOf(regPreCond))
}
