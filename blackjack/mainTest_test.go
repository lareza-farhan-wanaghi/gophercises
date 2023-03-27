package blackjack

import (
	"os"
	"testing"
)

type TestTable struct {
	retrieveCard    map[string]string
	rateCardsHasAce map[string]string
	startRound      map[string]string
}

var testTable TestTable

// populateRetrieveCard populates test cases for the retrieveCard function of the gameManager
func (tt *TestTable) populateRetrieveCard() {
	testCases := map[string]string{}
	testCases["a,b,c,d,e,1"] = "#r,1,e,#l,4,a,b,c,d"
	testCases["a,b,c,d,e,4"] = "#r,4,e,d,c,b,#l,1,a"
	testCases["a,b,2"] = "#r,2,b,a,#l,0"
	testCases["a,1"] = "#r,1,a"
	testCases["a,b,3"] = "#r,3,b,a,b,#l,1,a"

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

// populateStartRound populates test cases for the startRound function of the gameManager
func (tt *TestTable) populateStartRound() {
	testCases := map[string]string{}
	testCases["1,3,2,11,1,10,0"] = "0,16"
	testCases["1,10,2,11,3,4"] = "1,7"
	testCases["1,4,3,11,2,10,1"] = "0,14"
	testCases["1,4,10,3,1,10,11"] = "1,14"
	testCases["1,10,1,3,10,4"] = "1,12"
	testCases["1,2,3"] = "0,15"
	testCases["2,2,3"] = "0,19"
	testCases["2,2,3"] = "1,18"
	testCases["2,4,3,11,2,10,12,1,11"] = "0,20"
	testCases["2,3,11,2,10,0,1"] = "1,21"
	testCases["2,0,11,10,2,4,10,1,10"] = "1,16"

	tt.startRound = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateRetrieveCard()
	tt.populateRateCardsHasAce()
	tt.populateStartRound()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()
	pd_TurnedOff = true

	exitVal := m.Run()
	os.Exit(exitVal)
}
