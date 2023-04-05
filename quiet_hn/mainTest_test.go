package quiethn

import (
	"os"
	"testing"
)

type TestTable struct {
	getTopStories []string
	refreshCache  []string
}

var testTable TestTable

// populateGetTopStories populates test cases for the GetTopStories function of the client struct
func (tt *TestTable) populateGetTopStories() {
	testCases := []string{"30", "5", "0", "101"}

	tt.getTopStories = testCases
}

// populateRefreshCache populates test cases for the refreshCache function of the client struct
func (tt *TestTable) populateRefreshCache() {
	testCases := []string{"5", "10", "65"}

	tt.refreshCache = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateGetTopStories()
	tt.populateRefreshCache()

	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
