package main

import "fmt"

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func main() {

	xs := []string{"Alok", "Tripathi", "Alice", "Alex", "Alok", "Tripathi"}
	set := UniqStr(xs)
	fmt.Println(set)

}
