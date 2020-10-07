package main

import (
	"fmt"
	"strconv"
)

func main() {

	cards := newDeck(10)
	cards.print()

	hands, remainingDeck := deal(cards, 5)

	fmt.Println("hands", hands)
	fmt.Println("remainingDeck", remainingDeck)
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// fiveSpadesCard := getSpadesCard(5)

	// fmt.Println(card)
	// fmt.Println(fiveSpadesCard)

	// setDeck()
}

func getSpadesCard(number int) string {
	return strconv.Itoa(number) + " of Spades"
}
