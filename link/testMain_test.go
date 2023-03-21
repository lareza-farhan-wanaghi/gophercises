package link

import (
	"os"
	"testing"
)

type TestTable struct {
	aTagFinder      map[string][]*ATag
	convertToString map[string]string
}

var testTable TestTable

// populateATagFinder populates test cases for the ATagFinder functions
func (tt *TestTable) populateATagFinder() {
	testCases := make(map[string][]*ATag)
	testCases["ex1.html"] = []*ATag{
		{
			Href: "/other-page",
			Text: "A link to another page",
		},
	}
	testCases["ex2.html"] = []*ATag{
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github !",
		},
	}
	testCases["ex3.html"] = []*ATag{
		{
			Href: "#",
			Text: "Login",
		},
		{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}
	testCases["ex4.html"] = []*ATag{
		{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}
	tt.aTagFinder = testCases
}

// populateConvertToString populates test cases for the convertToString function
func (tt *TestTable) populateConvertToString() {
	testCases := make(map[string]string)
	testCases["Check me out on twitter "] = "Check me out on twitter"
	testCases[" Check me out on twitter"] = "Check me out on twitter"
	testCases[" Check me out on twitter "] = "Check me out on twitter"
	testCases["Check me out on twitter\n"] = "Check me out on twitter"
	testCases["Check me out on \n"] = "Check me out on"
	tt.convertToString = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateATagFinder()
	tt.populateConvertToString()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
