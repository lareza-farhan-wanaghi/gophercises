package main

import "github.com/lareza-farhan-wanaghi/gophercises/blackjack"

// main provides the entry point of the app
func main() {
	players := []*blackjack.Player{
		{Name: "Farhan", IsActive: true},
		{Name: "Comp", IsActive: false},
	}
	gm := blackjack.NewGameManager(players...)
	gm.Play()
}
