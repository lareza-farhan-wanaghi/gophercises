package blackjack

import (
	"bytes"
	"fmt"
	"time"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// rateCardsHasAce returns the score of the cards based on the blackjack rules and whether the card collection has an ace
func rateCardsHasAce(cards []*deck.Card) (int, bool) {
	aces := 0
	score := 0
	for i := 0; i < len(cards); i++ {
		val := cards[i].Id - cards[i].Id/100*100
		if val == 0 {
			aces += 1
		} else if val > 9 {
			score += 10
		} else {
			score += val + 1
		}
	}

	for i := 0; i < aces; i++ {
		if score <= 10 && (score == 10 || i == aces-1) {
			score += 11
		} else {
			score += 1
		}
	}

	return score, aces > 0
}

// rateCards returns the score of the cards based on the blackjack rules
func rateCards(cards ...*deck.Card) int {
	score, _ := rateCardsHasAce(cards)
	return score
}

// rateCards returns the score of the cards based on the blackjack rules
func cardString(cards ...*deck.Card) string {
	var sb bytes.Buffer
	for i, c := range cards {
		if i != len(cards)-1 {
			sb.WriteString(fmt.Sprintf("%s, ", c.Name))
		} else {
			sb.WriteString(fmt.Sprintf("and %s", c.Name))
		}
	}
	return sb.String()
}

const pd_fast = 100
const pd_long = 1000

var pd_TurnedOff bool

// printfDelay prints the string and sleeps the program for the given time afterward
func printfDelay(format string, delayTime int, args ...any) {
	if pd_TurnedOff {
		return
	}

	fmt.Printf(format, args...)
	time.Sleep(time.Duration(delayTime) * time.Millisecond)
}
