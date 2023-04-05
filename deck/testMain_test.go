package deck

import (
	"os"
	"testing"
)

type TestTable struct {
	newSuitDeck map[string][]func([]*Card) []*Card
}

var testTable TestTable

// populateDeck populates test cases for the NewSuitDeck function
func (tt *TestTable) populateDeck() {
	testCases := map[string][]func([]*Card) []*Card{}
	testCases["test1.txt"] = []func([]*Card) []*Card{
		SortAscendingCardId(),
	}
	testCases["test2.txt"] = []func([]*Card) []*Card{
		AppendSuitJokers(1),
		ExcludeSuitIds([]int{0, 100, 212, 312}),
		SortAscendingCardId(),
	}
	testCases["test3.txt"] = []func([]*Card) []*Card{
		AppendSuitJokers(1),
		ExcludeSuitIds([]int{5}),
		ConcatCards([]*Card{{6, "Spade Seven"}}),
		SortAscendingCardId(),
	}
	testCases["test4.txt"] = []func([]*Card) []*Card{
		AppendSuitJokers(2),
		ConcatCards(
			[]*Card{
				{0, "Spade Ace"},
				{0, "Spade Ace"},
				{1, "Spade Two"},
				{5, "Spade Six"},
				{5, "Spade Six"},
				{7, "Spade Eight"},
				{101, "Diamond Two"},
				{103, "Diamond Four"},
				{210, "Club Jack"},
				{212, "Club King"},
				{302, "Heart Three"},
			},
		),
		ExcludeSuitIds([]int{5, 7, 100, 210, 211, 300, 301, 400}),
		SortAscendingCardId(),
	}

	tt.newSuitDeck = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateDeck()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
