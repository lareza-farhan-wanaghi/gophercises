package transform

import (
	"os"
	"testing"
)

type TestTable struct {
	primitive []string
}

var testTable TestTable

// populatePrimitive populates test cases for the Primitive function
func (tt *TestTable) populatePrimitive() {
	testCases := []string{
		"images/test/below1MB.png,0,100",
		"images/test/below1MB.png,1,167",
		"images/test/about1MB.png,2,75",
		"images/test/about1MB.png,2,200",
		"images/test/above1MB.png,4,25",
		"images/test/above1MB.png,5,50",
	}

	tt.primitive = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populatePrimitive()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
