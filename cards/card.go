package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type cardTypes [4]string

func getCardTypes() cardTypes {
	cardTypes := cardTypes{"Hearts", "Spades", "Diamonds", "Stars"}
	return cardTypes
}

func getCrad() string {
	cardsTypes := getCardTypes()
	cardType := cardsTypes[rand.Intn(4)]
	fmt.Println("cardType", cardType)
	cardNumber := rand.Intn(10)
	fmt.Println("cardNumber", cardNumber)

	if cardNumber == 1 {
		return "Ace of " + cardType
	}
	card := strconv.Itoa(cardNumber) + " of " + cardType
	fmt.Println("card", card)
	return card
}
