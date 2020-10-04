package main

import "fmt"

type deck []string

// adding a new receiver
func (d deck) print() {
	fmt.Println("Deck cards:")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
