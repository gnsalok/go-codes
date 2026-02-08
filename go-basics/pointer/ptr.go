package main

import (
	"fmt"
)

func main() {
	bookSl := []string{
		"DSA1",
		"C",
		"C++",
	}

	AddBook("Hindi", bookSl)

	fmt.Println("In main", bookSl)

}

func AddBook(bookName string, bookSl []string) {

}
