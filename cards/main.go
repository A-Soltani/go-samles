package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	fiveSpadesCard := getSpadesCard(5)

	fmt.Print(card, "\n")
	fmt.Print(fiveSpadesCard)
}

func getSpadesCard(number int) string {
	return strconv.Itoa(number) + " of Spades"
}
