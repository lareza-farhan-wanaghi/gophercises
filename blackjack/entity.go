package blackjack

import (
	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

type entity interface {
	reset()
	enterPhase(gm *GameManager, args ...interface{})
	addCard(gm *GameManager, card *deck.Card)
	string() string
	getCards() []*deck.Card
}
