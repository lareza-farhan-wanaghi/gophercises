package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/blackjack"
	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// main provides the entry point of the app
func main() {
	compFlag := flag.Int("c", 1, "Specifies the number of AIs included in the game")
	flag.Parse()

	players := []*blackjack.Player{}
	for i := 0; i < *compFlag; i++ {
		player := &blackjack.Player{Name: fmt.Sprintf("Computer%d", i+1), IsActive: false}
		players = append(players, player)
	}

	for _, arg := range os.Args[3:] {
		player := &blackjack.Player{Name: arg, IsActive: true}
		players = append(players, player)
	}

	gm := blackjack.NewGameManager(deck.NewSuitDeck(deck.Shuffle()), players...)
	gm.Play()
}
