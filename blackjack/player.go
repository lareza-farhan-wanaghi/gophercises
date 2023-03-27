package blackjack

import (
	"bufio"
	"os"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
)

type Player struct {
	Name     string
	IsActive bool
	cards    []*deck.Card
	id       int
}

// enterPhase handles every phase of the game
func (p *Player) enterPhase(gm *GameManager, args ...interface{}) {
	printfDelay("\n### %s's Turn ###\n", pd_fast, p.Name)

	score := rateCards(p.cards...)
	printfDelay("Current cards are %s, and the score is %d\n", pd_fast, cardString(p.cards...), score)

	if p.IsActive {
		input := bufio.NewScanner(os.Stdin)
		for {
			printfDelay("Please press one of the options below to continue:\n", pd_fast)
			printfDelay("1. Hit\n", pd_fast)
			printfDelay("2. Stand\n", pd_fast)
			input.Scan()
			inputVal := input.Text()
			if inputVal == "1" {
				p.hit(gm)
				break
			} else if inputVal == "2" {
				p.stand()
				break
			} else {
				printfDelay("You entered an invalid option\n", pd_fast)
			}
		}
	} else {
		if score <= 10 {
			p.hit(gm)
		} else if score <= 15 && rateCards(gm.entities[len(gm.entities)-1].getCards()[1]) > 7 {
			p.hit(gm)
		} else {
			p.stand()
		}
	}

	printfDelay("### %s's End Turn\n", pd_long, p.Name)
}

// hit simulates the hit action of the game, which accepts the card offer
func (p *Player) hit(gm *GameManager) {
	printfDelay("%s chose to hit\n", pd_long, p.Name)
	gm.retrieveCard(p.id)
}

// hit simulates the stand action of the game, which rejects the card offer
func (p *Player) stand() {
	printfDelay("%s chose to stand\n", pd_long, p.Name)
}

// addCard adds the card to the list of the current cards held
func (p *Player) addCard(gm *GameManager, card *deck.Card) {
	p.cards = append(p.cards, card)
	score := rateCards(p.cards...)

	printfDelay("%s added a card, it's %s. Current score is %d\n", pd_fast, p.string(), card.Name, score)
	if score == 21 {
		gm.instantWin(p.id)
	} else if score > 21 {
		gm.eliminate(p.id)
	}
}

// string return the string representation of this object
func (p *Player) string() string {
	return p.Name
}

// getCards returns the cards help by this object
func (p *Player) getCards() []*deck.Card {
	return p.cards
}

// reset sets the object to its default stat
func (p *Player) reset() {
	p.cards = []*deck.Card{}
}
