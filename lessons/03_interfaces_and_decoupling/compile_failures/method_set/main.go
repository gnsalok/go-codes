// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This example intentionally fails to demonstrate method-set rules.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
	showUser()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u user) showUser() {
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
	u := user{"Bill", "bill@email.com"}

	// user does not implement notifier because notify has a pointer receiver.
	// Passing &u is the corrected form.
	sendNotification(u)
}
