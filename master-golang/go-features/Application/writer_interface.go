package main

import (
	"io/ioutil"
)

func main() {

	data := []byte("This is sample text")
	err := ioutil.WriteFile("sample.txt", data, 0777)

}
