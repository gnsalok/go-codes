package main

import (
	"errors"
	"log"
)

type ownError struct{}

var (
	ErrOwnBr = errors.New("Bad Request")
)

func main() {

	if err := webCall(true); err == ErrOwnBr {
		log.Print("There is a problem")

	}
	log.Println("Life is Good.")

}

func webCall(b bool) error {
	if b {
		return errors.New("Bad Request")
	}
	return nil
}
