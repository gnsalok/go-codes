// Interface naming or declaration
package main

// interface should named like, name + (er) in the end
// Writer, Reader etc

// Define interface like exactly what to do (single responsibility) and group interfaces into one.

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

type Storer interface {
	Getter
	Putter
}

func main() {

}
