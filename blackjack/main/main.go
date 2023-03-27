package main

import (
	"github.com/lareza-farhan-wanaghi/gophercises/blackjack"
	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// main provides the entry point of the app
func main() {
	players := []*blackjack.Player{
		{Name: "Farhan", IsActive: true},
		{Name: "Comp", IsActive: false},
	}
	gm := blackjack.NewGameManager(deck.NewSuitDeck(deck.Shuffle()), players...)
	gm.Play()
}
