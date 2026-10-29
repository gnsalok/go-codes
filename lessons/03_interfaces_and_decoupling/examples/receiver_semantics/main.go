package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"bill", "bill@gmail.com"}

	// bill is addressable, so the compiler can take &bill for this call.
	bill.changeEmail("bill@tgmail.com")
	bill.notify()

	jon := &user{"jon", "jon@gmail.com"}
	jon.changeEmail("jon@tgmail.com")
	jon.notify()

	users := []user{
		{"at", "at@gmail.com"},
		{"arik", "arik@gmail.com"},
	}

	// A range value is a copy. Mutating it does not update the slice element.
	for _, u := range users {
		u.changeEmail("it@doesntmatter.com")
	}

	for _, u := range users {
		fmt.Println(u.email)
	}

	// Indexing addresses the actual element.
	for i := range users {
		users[i].changeEmail("it@doesntmatter.com")
	}

	for _, u := range users {
		fmt.Println(u.email)
	}
}
