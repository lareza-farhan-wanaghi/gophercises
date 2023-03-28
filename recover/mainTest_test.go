package recover

import (
	"os"
	"testing"
)

type TestTable struct {
	getRequest map[string]string
}

var testTable TestTable

// populateGetRequest populates test cases simulating web get requests
func (tt *TestTable) populateGetRequest() {
	testCases := map[string]string{}
	testCases["/"] = "200,Hello!"
	testCases["/panic/"] = "500,Something went wrong"
	tt.getRequest = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateGetRequest()

	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
