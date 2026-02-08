package main

import "fmt"

func main() {
	name := "Alok Tripathi"
	course := "CS101"

	fmt.Println("Name:", name, "Course:", course)

	updateCourse(&course)
	fmt.Println("Course in Main", course)

}

func updateCourse(course *string) string {
	*course = "CS100" // Updating the actual value of course variable in main
	fmt.Println("\nCourse inside function", *course)
	return *course
}
