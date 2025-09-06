package main

import (
	"fmt"
	"sync"
	"time"
)

type pair struct {
	row, col int
}

const length = 2

var start time.Time
var result [length][length]int

func main() {
	const threadlength = 1

	pairs := make(chan pair, 100)
	var wg sync.WaitGroup

	//Change this input to read from file
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{5, 6}, {7, 8}}

	fmt.Println("Matrix A")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(a[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

	fmt.Println("Matrix B")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(b[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

	wg.Add(threadlength)

	for i := 0; i < threadlength; i++ {
		go addMatrix(pairs, &a, &b, &result, &wg)
	}

	start = time.Now()

	// assigning row and coloums value to struct
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			pairs <- pair{row: i, col: j}
		}
	}

	close(pairs)

	wg.Wait()

	fmt.Print("---------------------------------------------\n")
	fmt.Println("Matrix (A + B) ")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(result[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)
}

// AddMatrix
func addMatrix(pairs chan pair, a, b, result *[length][length]int, wg *sync.WaitGroup) {
	for {
		pair, ok := <-pairs
		if !ok {
			break
		}

		result[pair.row][pair.col] = 0

		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				result[i][j] = a[i][j] + b[i][j]
			}
		}
	}
	wg.Done()
}
