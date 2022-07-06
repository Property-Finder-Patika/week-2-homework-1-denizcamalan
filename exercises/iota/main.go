package main

import "fmt"

type Season int32

const (
	Summer Season = 1 + iota
	Autumn
	Winter
	Spring
)

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func main() {
	printSeason(Summer)
}
