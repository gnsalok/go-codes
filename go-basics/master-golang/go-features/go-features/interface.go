package main

import "fmt"

type Guitarist interface {
    // PlayGuitar prints out "Playing Guitar"
    // to the terminal
    PlayGuitar()
}

type BaseGuitarist struct {
    Name string
}

type AcousticGuitarist struct {
    Name string
}

//can pass b *BaseGuitarist if we want to make global change.
func (b BaseGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
    var player BaseGuitarist
    player.Name = "Paul"
    player.PlayGuitar()

    var player2 AcousticGuitarist
    player2.Name = "Ringo"
    player2.PlayGuitar()
}