package recover

import (
	"os"
	"testing"
)

type TestTable struct {
	getRequest map[string]string
	parseTrace map[string]string
}

var testTable TestTable

// populateGetRequest populates test cases simulating web get requests
func (tt *TestTable) populateGetRequest() {
	testCases := map[string]string{}
	testCases["/"] = "200,Hello!"
	testCases["/panic/"] = "500,Something went wrong"
	tt.getRequest = testCases
}

// populateParseTrace populates test cases fo the parseTrace function
func (tt *TestTable) populateParseTrace() {
	testCases := map[string]string{}
	testCases["/usr/blah.go:24"] = "<a href='/debug/usr/blah.go?id=24'>/usr/blah.go:24</a>"
	testCases["/usr/blah.go:22"] = "<a href='/debug/usr/blah.go?id=22'>/usr/blah.go:22</a>"
	testCases["/usr/blah.go:8 +0x65"] = "<a href='/debug/usr/blah.go?id=8'>/usr/blah.go:8 +0x65</a>"
	testCases["/usr:8 +0x65\n/hey:9"] = "<a href='/debug/usr?id=8'>/usr:8 +0x65</a>\n<a href='/debug/hey?id=9'>/hey:9</a>"
	testCases["/usr:8 +0x65\n/hey:9\n/back/back:79"] = "<a href='/debug/usr?id=8'>/usr:8 +0x65</a>\n<a href='/debug/hey?id=9'>/hey:9</a>\n<a href='/debug/back/back?id=79'>/back/back:79</a>"

	tt.parseTrace = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateGetRequest()
	tt.populateParseTrace()

	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
