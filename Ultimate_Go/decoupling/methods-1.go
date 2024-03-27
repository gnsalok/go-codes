package main

import "fmt"

/*
   Behaviour
	 |
   Decouple
	 |
   Concrete
	 |
   Data

 Follow -> Bottom to Top
---

Question?
When to use  value receiver/method?
When to use ptr receiver/method?

Answer :

-> Read markdown file
__

Note  (Best Practice):-
- Value semantic shares COPY, whereas Pointer semantic giving access.
- Don't change the semantics; From value semantic to ptr semantic, that can cause issue.
- Semantic consistency is Everything.

*/

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

	// Value of type user can be used to call methods.
	// Declared with both value and pointer receiver.
	bill := user{"bill", "bill@gmail.com"}

	// changeEmail is of Pointer semantic ; yes, it will compile :)
	// Go only cares about DATA. It doesn't care about if it is Value or Pointer.
	// All it is care about User value in some form.
	bill.changeEmail("bill@tgmail.com") // (&bill).ChangeEmail
	bill.notify()

	// Pointer of type user can be used to call methods.
	// Declared with both value and pointer receiver.
	jon := &user{"jon", "jon@gmail.com"}

	// Avoid doing this ; No need to take copy of the value where pointer points to.
	// Always maintain Value semantic
	jon.changeEmail("jon@tgmail.com")
	jon.notify()

	// Create slice of users
	users := []user{
		{"at", "at@gmail.com"},
		{"arik", "arik@gmail.com"},
	}

	// iterate over the slice of user (copy of users) and changeEmail is ptr usage ptr semantic.
	// Mixing semantic NOT GOOD!!
	// Actual email in sliceis not going to change but copy is going to change
	for _, u := range users {
		u.changeEmail("it@doesntmatter.com")
	}

	// email will not change here
	for _, u := range users {
		fmt.Println(u.email)
	}

	// HACKING - you can change the email as index i usage ptr semantic
	for i, _ := range users {
		users[i].changeEmail("it@doesntmatter.com")
	}

	for _, u := range users {
		fmt.Println(u.email)
	}
}
