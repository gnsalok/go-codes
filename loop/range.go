package main

import "fmt"

func main() {

	courseList := []string{"Course1", "Course2", "Course3"}

	for i, v := range courseList {
		fmt.Println(i, v)
	}

}
