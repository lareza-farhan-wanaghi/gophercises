package blackjack

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

// TestRetrieveCard tests the retrieveCard function of the gameManager struct
func TestRetrieveCard(t *testing.T) {
	for k, v := range testTable.retrieveCard {
		splits := strings.Split(k, ",")
		cards := []*deck.Card{}
		for _, s := range splits[:len(splits)-1] {
			cards = append(cards, &deck.Card{Id: 0, Name: s})
		}

		numOfRetrieving, err := strconv.Atoi(splits[len(splits)-1])
		if err != nil {
			t.Fatal(err)
		}

		gm := NewGameManager(cards)
		for i := 0; i < numOfRetrieving; i++ {
			gm.retrieveCard(0)
		}

		recievedCards := gm.entities[0].getCards()
		splits = strings.Split(v, ",")
		for i, c := 0, ""; i < len(splits); {
			if c == "" {
				c = splits[i]
				i++
			} else {
				switch c {
				case "#r":
					m, err := strconv.Atoi(splits[i])
					if err != nil {
						t.Fatal(err)
					}
					i += 1

					if m != len(recievedCards) {
						t.Fatalf("expected length %d but got %d. k: %s v: %s", m, len(recievedCards), k, v)
					}

					for j := 0; j < m; j++ {
						if recievedCards[j].Name != splits[i+j] {
							t.Fatalf("expected %s but got %s. k: %s v: %s", splits[i+j], recievedCards[j].Name, k, v)
						}
					}
					i += m
				case "#l":
					m, err := strconv.Atoi(splits[i])
					if err != nil {
						t.Fatal(err)
					}
					i += 1

					if m != len(gm.cards) {
						t.Fatalf("expected length %d but got %d. k: %s v: %s", m, len(gm.cards), k, v)
					}

					for j := 0; j < m; j++ {
						if gm.cards[j].Name != splits[i+j] {
							t.Fatalf("expected %s but got %s. k: %s v: %s", splits[i+j], gm.cards[j].Name, k, v)
						}
					}
					i += m
				}
				c = ""
			}
		}
	}
}

// TestStartRound tests the startRound function of the gameManager struct
func TestStartRound(t *testing.T) {
	for k, v := range testTable.startRound {
		splits := strings.Split(k, ",")
		numOfPlayer, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		players := []*Player{}
		for i := 0; i < numOfPlayer; i++ {
			players = append(players, &Player{Name: fmt.Sprintf("p%d", i), IsActive: false})
		}

		cards := []*deck.Card{}
		for i := 1; i < len(splits); i++ {
			id, err := strconv.Atoi(splits[i])
			if err != nil {
				t.Fatal(err)
			}
			cards = append(cards, &deck.Card{Id: id})
		}

		gm := NewGameManager(cards, players...)
		gm.startRound()

		splits = strings.Split(v, ",")
		expectedId, err := strconv.Atoi(splits[0])
		if err != nil {
			t.Fatal(err)
		}

		if gm.winner != expectedId {
			t.Fatalf("expected %v but got %v. k: %s v: %s", expectedId, gm.winner, k, v)
		}

		expectedScore, err := strconv.Atoi(splits[1])
		if err != nil {
			t.Fatal(err)
		}

		score := rateCards(gm.entities[gm.winner].getCards()...)
		if score != expectedScore {
			t.Fatalf("expected %v but got %v. k: %s v: %s", expectedScore, score, k, v)
		}
	}
}
