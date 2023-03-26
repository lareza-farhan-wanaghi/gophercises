package blackjack

import (
	"os"
	"testing"
)

type TestTable struct {
	retrieveCard    map[string]string
	rateCardsHasAce map[string]string
}

var testTable TestTable

// populateRetrieveCard populates test cases for the retrieveCard function of the gameManager
func (tt *TestTable) populateRetrieveCard() {
	testCases := map[string]string{}
	testCases["a,b,c,d,e,1"] = "#r,1,e,#l,4,a,b,c,d"
	testCases["a,b,c,d,e,4"] = "#r,4,e,d,c,b,#l,1,a"
	testCases["a,b,2"] = "#r,2,b,a,#l,0"
	testCases["a,1"] = "#r,1,a"

	tt.retrieveCard = testCases
}

// populateRateCardsHasAce populates test cases for the rateCardsHasAce function of the gameManager
func (tt *TestTable) populateRateCardsHasAce() {
	testCases := map[string]string{}
	testCases["2,3,4"] = "12,0"
	testCases["0,2,3"] = "18,1"
	testCases["0,8,9"] = "20,1"
	testCases["0,10,11"] = "21,1"
	testCases["0,10,11,10"] = "31,1"
	testCases["10,11,10"] = "30,0"
	testCases["2,11,10"] = "23,0"

	tt.rateCardsHasAce = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateRetrieveCard()
	tt.populateRateCardsHasAce()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
