// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
	showuser()
}

// user defines a user in the program.
type userstr struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *userstr) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// Additional Context
// ISSUE : Do not try to mix the value and pointer semantic of struct while implementing methods
// It's not recommended; Wht? Look issue mixing interface semantic
// Refer : https://www.ardanlabs.com/blog/2017/07/interface-semantics.html
func (u userstr) showuser() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

func main() {

	// Create a value of type User and send a notification.
	u := userstr{"Bill", "bill@email.com"}

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	// Error
	// You can only pass address of u, as point struct method do not accept COPY/ value struct.
	// Make it `&u` it will work.
	sendNotification(u)

	// ./example1.go:36: cannot use u (type user) as type notifier in argument to sendNotification:
	//   user does not implement notifier (notify method has pointer receiver)
}
