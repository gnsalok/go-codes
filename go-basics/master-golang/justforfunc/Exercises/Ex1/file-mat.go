package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	var a [length][length]int
	var b [length][length]int

	// Reading matrices values from file1 and file2 respe...
	mat1 := readFile("src/github.com/gnsalok/master-golang/go-features/justforfunc/ovn-test/file1.txt")
	mat2 := readFile("src/github.com/gnsalok/master-golang/go-features/justforfunc/ovn-test/file2.txt")

	// Converting 1d array into 2d array
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			a[i][j] = mat1[(j*length)+i]
			b[i][j] = mat2[(j*length)+i]
		}
	}

	// Printing Matrix A
	fmt.Println("---------------Matrix A-----------------")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(a[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

	// Priting Matrix B
	fmt.Println("---------------Matrix B-----------------")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(b[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

	wg.Add(threadlength)

	// Putting pairs into the Channel
	for i := 0; i < threadlength; i++ {
		go AddMatrix(pairs, &a, &b, &result, &wg)
	}

	start = time.Now()

	// Assigning row and coloums values into struct
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			pairs <- pair{row: i, col: j}
		}
	}

	close(pairs)
	wg.Wait()

	//Priting Result of addition
	fmt.Println("----------------Matrix (A + B)-----------------")
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(result[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

	//Write output into the file
	writeLines(result, "src/github.com/gnsalok/master-golang/go-features/justforfunc/ovn-test/output.txt")

	// Printing elapsed time
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)
}

// AddMatrix adding 2 matrix
func AddMatrix(pairs chan pair, a, b, result *[length][length]int, wg *sync.WaitGroup) {
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

// Reading from the file
func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {

		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
	return
}

// Wrting addition output into the output.txt file
func writeLines(lines [length][length]int, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Println("Data written into output.txt file successfully!")
	return w.Flush()

}
