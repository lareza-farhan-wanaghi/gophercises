package blackjack

import (
	"strconv"
	"strings"
	"testing"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// TestRateCardsHasAce tests the rateCardsHasAce function
func TestRateCardsHasAce(t *testing.T) {
	for k, v := range testTable.rateCardsHasAce {
		cards := []*deck.Card{}
		for _, s := range strings.Split(k, ",") {
			id, err := strconv.Atoi(s)
			if err != nil {
				t.Fatal(err)
			}
			cards = append(cards, &deck.Card{Id: id})
		}

		splits := strings.Split(v, ",")
		score, hasAce := rateCardsHasAce(cards)
		expectedScore, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		if score != expectedScore {
			t.Fatalf("expected %d but got %d. k: %s v:%s", expectedScore, score, k, v)
		}

		boolInt, err := strconv.Atoi(splits[1])
		if err != nil {
			t.Fatal(err)
		}
		expectedHasAce := boolInt == 1

		if hasAce != expectedHasAce {
			t.Fatalf("expected %v but got %v. k: %s v: %s", expectedHasAce, hasAce, k, v)
		}
	}
}
