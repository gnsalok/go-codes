// function that panic always include Must as prefix

package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1"
	value := MustParseIntFromString(str)
	fmt.Println(value)
	// fmt.Println(reflect.TypeOf(str))

}

func MustParseIntFromString(s string) int64 {
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}
