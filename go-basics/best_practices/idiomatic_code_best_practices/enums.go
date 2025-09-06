// construct like enums (enums are not supported in go)

package main

import "fmt"

type Season int64

// always append type before name
const (
	SeasonSummer Season = iota
	SeasonAutumn
	SeasonWinter
	SeasonSpring
)

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func main() {
	var i Season
	i = 10
	printSeason(i)

}
