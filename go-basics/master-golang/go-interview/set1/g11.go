package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
}
