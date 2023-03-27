package blackjack

import (
	"bufio"
	"os"
	"strings"

	"github.com/lareza-farhan-wanaghi/gophercises/deck"
	"golang.org/x/tools/container/intsets"
)

type GameManager struct {
	cards          []*deck.Card
	cardsMaxLength int
	entities       []entity
	eliminated     map[int]struct{}
	winner         int
}

// Play enters the main inner-loop of the game
func (gm *GameManager) Play() {
	wins := make([]int, len(gm.entities))
	input := bufio.NewScanner(os.Stdin)
	for {
		printfDelay("%s\n", pd_fast, strings.Repeat("-", 30))
		printfDelay("Current statistic:\n", pd_fast)
		for i := 0; i < len(wins); i++ {
			printfDelay("%d. %s: %d wins\n", pd_fast, i+1, gm.entities[i].string(), wins[i])
		}
		printfDelay("Press any keys to start the game or N to end the game\n", pd_fast)
		input.Scan()
		inputVal := input.Text()
		if strings.ToLower(inputVal) == "n" {
			break
		}
		printfDelay("### Starting a round! ###\n", pd_long)
		gm.startRound()

		printfDelay("\n### GAME OVER ###\n", pd_long)
		printfDelay("The round is finished with %s as the winner\n\n", pd_fast, gm.entities[gm.winner].string())
		wins[gm.winner] += 1
		printfDelay("Latest entities' cards:\n", pd_fast)
		for i, e := range gm.entities {
			eliminatedInfo := ""
			if _, ok := gm.eliminated[i]; ok {
				eliminatedInfo = " (eliminated)"
			}
			printfDelay("%d. %s%s: %s. The score is %d\n",
				pd_fast, i+1, e.string(), eliminatedInfo, cardString(e.getCards()...), rateCards(e.getCards()...))
		}
		printfDelay("\n", pd_long)
	}
}

// Play starts the game round and determines the winner for that round
func (gm *GameManager) startRound() {
	gm.cards = gm.cards[:gm.cardsMaxLength]
	gm.eliminated = map[int]struct{}{}
	gm.winner = -1

	for _, e := range gm.entities {
		e.reset()
	}

	for i := 0; i < 2; i++ {
		for j := range gm.entities {
			gm.retrieveCard(j)
		}
	}

	if gm.winner >= 0 {
		return
	} else if len(gm.eliminated) >= len(gm.entities)-1 {
		gm.winner = len(gm.entities) - 1
		return
	}

	for i := 0; i < 2; i++ {
		for j, e := range gm.entities {
			if _, ok := gm.eliminated[j]; ok {
				continue
			}

			e.enterPhase(gm, i)
			if gm.winner >= 0 {
				return
			} else if len(gm.eliminated) >= len(gm.entities)-1 {
				gm.winner = len(gm.entities) - 1
				return
			}
		}
	}

	highestScore := intsets.MinInt
	for i, e := range gm.entities[:len(gm.entities)-1] {
		score := rateCards(e.getCards()...)
		_, ok := gm.eliminated[i]
		if !ok && score <= 21 && score > highestScore {
			gm.winner = i
			highestScore = score
		}
	}
}

// retrieveCard retrieve a card from the deck and adds it to the entity
func (gm *GameManager) retrieveCard(id int) *deck.Card {
	if len(gm.cards) == 0 {
		gm.cards = gm.cards[:gm.cardsMaxLength]
	}
	card := gm.cards[len(gm.cards)-1]
	gm.cards = gm.cards[:len(gm.cards)-1]
	gm.entities[id].addCard(gm, card)
	return card
}

// instantWin forces every player to hit a card from the deck
func (gm *GameManager) playerHits() {
	printfDelay("Due to a dealt from the dealer, all player must hit a card\n", pd_fast)
	for i := 0; i < len(gm.entities)-1; i++ {
		_, ok := gm.eliminated[i]
		if !ok {
			gm.retrieveCard(i)
		}
	}
}

// eliminate eliminates the entity for this round
func (gm *GameManager) eliminate(id int) {
	gm.eliminated[id] = struct{}{}
	printfDelay("Due to a bust, %s is eliminated\n", pd_fast, gm.entities[id].string())
}

// instantWin turns the entity as the winner for this round
func (gm *GameManager) instantWin(id int) {
	gm.winner = id
	printfDelay("Due to a backjack21, %s wins the round\n", pd_fast, gm.entities[id].string())
}

// NewGameManager creates a game manager with the players included
func NewGameManager(cards []*deck.Card, players ...*Player) *GameManager {
	entities := []entity{}
	for i, p := range players {
		p.id = i
		entities = append(entities, entity(p))
	}
	entities = append(entities, entity(&dealer{id: len(entities)}))

	return &GameManager{
		entities:       entities,
		cards:          cards,
		cardsMaxLength: len(cards),
	}
}
