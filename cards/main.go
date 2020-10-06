package main

import (
	"strconv"
)

func main() {

	setDeck()

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

func setDeck() {
	cards := newDeck(5)
	// cards = append(cards, "6 of Spades")

	cards.print()
}
