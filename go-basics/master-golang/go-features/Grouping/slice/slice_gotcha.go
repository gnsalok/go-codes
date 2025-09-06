/*



 */

package main

import (
	"fmt"
)

func main() {

	s := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(s)
	fmt.Println(s)
	fmt.Printf("Address of slice %p\n", &s)

}

func updateSlice(s []string) {
	s[0] = "Hello"
	fmt.Printf("Address of slice %p\n", &s)
}
