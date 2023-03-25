// Package deck creates a card slice that represent a deck of cards for a given settings
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Card represents a card for general card games.
type Card struct {
	Id   int
	Name string
}

// String stringifies contents of the object
func (c *Card) String() string {
	return fmt.Sprintf("%v - %v", c.Id, c.Name)
}

// SortAscendingCardId returns a card sorting function that sorts ascendingly the cards by their ids
func SortAscendingCardId() func(prevCards []*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		sort.Slice(prevCards, func(i, j int) bool {
			return prevCards[i].Id < prevCards[j].Id
		})
		return prevCards
	}
}

// SortDescendingCardId returns a card sorting function that sorts descendingly the cards by their ids
func SortIdDescending() func(prevCards []*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		sort.Slice(prevCards, func(i, j int) bool {
			return prevCards[i].Id > prevCards[j].Id
		})
		return prevCards
	}
}

// Shuffle returns a function that shuffles the card collection
func Shuffle() func([]*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		r1.Shuffle(len(prevCards), func(i, j int) {
			tmp := prevCards[i]
			prevCards[i] = prevCards[j]
			prevCards[j] = tmp
		})
		return prevCards
	}
}

// ConcatCards concats arbitrary two card collections
func ConcatCards(cardsToCancat []*Card) func(prevCards []*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		result := append(prevCards, cardsToCancat...)
		return result
	}
}

// NewSuitDeck return suit deck cards.
//
// Arguments:
//   - funcOps: specifies specific options that will be executed sequentially
func NewSuitDeck(funcOps ...func([]*Card) []*Card) []*Card {
	result := getSuitStandartCard()
	for _, f := range funcOps {
		result = f(result)
	}
	return result
}
