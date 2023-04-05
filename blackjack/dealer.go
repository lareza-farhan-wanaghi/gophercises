package blackjack

import (
	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

type dealer struct {
	cards []*deck.Card
	id    int
}

// enterPhase handles every phase of the game
func (d *dealer) enterPhase(gm *GameManager, args ...interface{}) {
	printfDelay("\n### %s's Turn ###\n", pd_fast, d.string())

	if args[0].(int) == 0 {
		printfDelay("%s skips its turn during this phase\n", pd_long, d.string())
	} else {
		printfDelay("%s reveal its cards. They are %s\n", pd_fast, d.string(), cardString(d.cards...))
		score, hasAce := rateCardsHasAce(d.cards)
		if score < 17 {
			printfDelay("The score is %d, which is below 17. Thus, all players must hit a card\n", pd_long, score)
			gm.playerHits()
		} else if score == 17 {
			if hasAce {
				printfDelay("The score is %d, and it's an soft-17. Thus, all players must hit a card\n", pd_long, score)
				gm.playerHits()
			} else {
				printfDelay("The score is %d, but it's not an soft-17. Thus, all players must stand with their cards\n", pd_long, score)
			}
		} else {
			printfDelay("The score is %d, which is above 17. Thus, all players must stand with their cards\n", pd_long, score)
		}
	}
	printfDelay("### Dealer's End Turn\n", pd_long)
}

// addCard adds the card to the list of the current cards held
func (d *dealer) addCard(gm *GameManager, card *deck.Card) {
	d.cards = append(d.cards, card)
	cardInfo := card.Name
	if len(d.cards) == 1 {
		cardInfo = "unrevealed"
	}
	printfDelay("%s added a card, it's %s\n", pd_fast, d.string(), cardInfo)
}

// string return the string representation of this object
func (d *dealer) string() string {
	return "Dealer"
}

// getCards returns the cards help by this object
func (p *dealer) getCards() []*deck.Card {
	return p.cards
}

// reset sets the object to its default stat
func (d *dealer) reset() {
	d.cards = []*deck.Card{}
}
