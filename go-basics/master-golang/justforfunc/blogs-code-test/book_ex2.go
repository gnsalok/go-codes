package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

func NewUser(id int, name, location string) *User {
	//	id++
	return &User{id, name, location}
}

func main() {
	u := NewUser(42, "Matt", "LA")
	fmt.Println(u.Greetings())
}
