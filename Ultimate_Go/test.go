package main

func main() {
	var sl []string
	sl = append(sl, "foo", "bar", "baz")
	for i, v := range sl {
		println(i, v)
	}

	newSl := make([]string, len(sl))
	copy(newSl, sl)
	for i, v := range newSl {
		println(i, v)
	}

	newSl = append(newSl[:1], newSl[2:]...)
	for i, v := range newSl {
		println(i, v)

	}
	println("After removal:")
	for i, v := range newSl {
		println(i, v)
	}

}
