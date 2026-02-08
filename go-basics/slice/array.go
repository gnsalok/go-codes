package main

import "fmt"

func main() {

	// Array is of fixed length
	var courseList [3]string
	courseList[0] = "course1"
	courseList[1] = "course2"
	courseList[2] = "course3"

	// courseList := []string{"Course1", "Course2", "Course3"}

	for i, v := range courseList {
		fmt.Println(i, v)
	}

}
