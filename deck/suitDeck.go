package deck

import (
	"fmt"
)

//go:generate stringer -type=SuitType
type SuitType int

const (
	Spade SuitType = iota
	Diamond
	Club
	Heart
	Joker
)

//go:generate stringer -type=SuitValue
type SuitValue int

const (
	Ace SuitValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// getSuitStandartCard returns all suit-collection standard cards
//
// The ids are specified by (100 * <suitType> + <suitValue>) calculations
// (for Jokers the suitValue values are 0).
// For example: id of 1 means an ace with a suit of Spade,
// 111 means a jack with a suit of Diamond, and 400 means a joker
//
// Meanwhile, the names are specified by "<val-value> <suit-value>" format
// (for jokers the format will be "<val-value>")
func getSuitStandartCard() []*Card {
	result := []*Card{}
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			id := j*100 + i
			cardName := getSuitCardName(SuitType(j), SuitValue(i))

			card := Card{
				Id:   id,
				Name: cardName,
			}
			result = append(result, &card)
		}
	}
	return result
}

// AppendSuitJokers appends suit joker cards to the card collection
func AppendSuitJokers(numOfJokers int) func([]*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		result := prevCards
		jokerId := 400

		for i := 0; i < numOfJokers; i++ {
			card := Card{
				Id:   jokerId,
				Name: "Joker",
			}
			result = append(result, &card)
		}
		return result
	}
}

// ExcludeSuitIds removes suit cards by their ids from the card collection
func ExcludeSuitIds(cardIdsToExclude []int) func([]*Card) []*Card {
	return func(prevCards []*Card) []*Card {
		excludedIdMap := map[int]struct{}{}
		for _, id := range cardIdsToExclude {
			excludedIdMap[id] = struct{}{}
		}

		result := []*Card{}
		for _, card := range prevCards {
			_, ok := excludedIdMap[card.Id]
			if !ok {
				result = append(result, card)
			}
		}
		return result
	}
}

// getSuitCardName gets the name of a suit card based on its atribute
func getSuitCardName(suit SuitType, val SuitValue) string {
	return fmt.Sprintf("%s %s", suit.String(), val.String())
}
