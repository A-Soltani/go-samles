package main

import (
	"fmt"
)

type deck []string

func newDeck(cardsCount int) deck {
	cards := deck{"Ace of Diamonds"} // this is a slice
	for i := 0; i < cardsCount; i++ {
		card := getCrad()
		// if !cards.deckContainsCard(card) {
		cards = append(cards, card)
		// }
	}

	return cards
}

func (d deck) deckContainsCard(card string) bool {
	for _, deckCard := range d {
		if card == deckCard {
			return false
		}
	}
	return true
}

// adding a new receiver
func (d deck) print() {
	fmt.Println("Deck cards:")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
