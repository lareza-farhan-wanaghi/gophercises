package main

import (
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// main provides the entry point of the app
func main() {
	cards := deck.NewSuitDeck(deck.SortAscendingCardId())
	for _, card := range cards {
		fmt.Printf("%v\n", card.String())
	}
}
