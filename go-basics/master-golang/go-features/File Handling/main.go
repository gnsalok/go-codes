package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//Reading the file
	// data, err := ioutil.ReadFile("myfile.data")

	// if err != nil{
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(data))

	//writing into the file
	data := []byte("This is sample text")

	err := ioutil.WriteFile("sample.txt", data, 0777)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data written successfully!")
	}
}
