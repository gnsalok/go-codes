// Implement File Handling in Go.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
1. Create a file
2. Write inside it
3. Print char inside file in console.
*/

func main() {

	//Create the file
	file, err := os.Create("Sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Writing inside the file
	file.WriteString("I'm Alok and I'm passionate about Software Engineering")

	//Extract string out of the file
	stream, err := ioutil.ReadFile("Sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	//reading text inside the file
	text := string(stream)
	fmt.Println(text)

}
