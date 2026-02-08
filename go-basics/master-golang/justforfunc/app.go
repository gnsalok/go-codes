package main

import (
	"fmt"
	"strconv"
)

func numtowords(val int) string {
	samarra := [10]string{"zero ", "one ", "two ", "three ", "four ", "five ", "six ", "seven ", "eight ", "nine "}
	return samarra[val]
}

func main() {
	argsinput := []string{"1", "2", "3", "4"}
	var output string
	var input string

	//argsinput = os.Args[1:]
	for j := 0; j < len(argsinput); j++ {
		input = argsinput[j]
		output = ""
		for i := 0; i < len(input); i++ {
			sample, err := strconv.Atoi(string(input[i]))
			if input[i] == ' ' && err != nil {
				fmt.Println(output)
				output = ""
			} else if sample < 10 && err == nil {
				output = output + numtowords(sample)
			} else {
				fmt.Println("Not a Number")
			}
		}
		fmt.Println(output)
		fmt.Println()
	}
}
