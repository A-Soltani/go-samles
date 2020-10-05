package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{"Ace of Diamonds"} // this is a slice
	return cards
}

// adding a new receiver
func (d deck) print() {
	fmt.Println("Deck cards:")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
