package hr1

import (
	"os"
	"testing"
)

type TestTable struct {
	camelcase    map[string]int32
	caesarChiper map[string]string
}

var testTable TestTable

// populateCamelcase populates test cases for the camelcase function
func (tt *TestTable) populateCamelcase() {
	testCases := make(map[string]int32)
	testCases[""] = 0
	testCases["Am"] = 0
	testCases["1am"] = 0
	testCases["hello"] = 1
	testCases["helloThere"] = 2
	testCases["anotherHelloThere"] = 3
	tt.camelcase = testCases
}

// populateCaesarChiper populates test cases for the caesarChiper function
func (tt *TestTable) populateCaesarChiper() {
	testCases := make(map[string]string)
	testCases["middle-Outz,2"] = "okffng-Qwvb"

	tt.caesarChiper = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateCamelcase()
	tt.populateCaesarChiper()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
