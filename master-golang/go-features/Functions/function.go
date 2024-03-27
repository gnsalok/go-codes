package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//added comment 

//receiver method it is not function, the idea here is the value of type(person) person has access to this method.
func (p *Person) update(n string, a int) {
	p.name = n
	p.age = a
	fmt.Println(p.name, p.age)
}


//Function
//func update(p Person){ }

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func main() {
	var p Person
	p.update("Alok", 21)

	mess := woo("Hi there!")
	fmt.Println(mess)

}
