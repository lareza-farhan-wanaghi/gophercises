package main

import (
	"flag"
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

var sortingAlgorithm = []func([]*deck.Card) []*deck.Card{
	deck.SortAscendingCardId(),
	deck.SortIdDescending(),
	deck.Shuffle(),
}

// main provides the entry point of the app
func main() {
	jokerFlag := flag.Int("j", 0, "Specifies the number of jokers included")
	sortFlag := flag.Int("s", 0, "Specifies the index of the sorting algorithm")
	flag.Parse()

	cards := deck.NewSuitDeck(deck.AppendSuitJokers(*jokerFlag), sortingAlgorithm[*sortFlag])

	for _, card := range cards {
		fmt.Printf("%v\n", card.String())
	}
}
