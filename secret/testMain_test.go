package secret

import (
	"os"
	"testing"
)

type TestTable struct {
	get    map[string]string
	getAll map[string]string
}

var testTable TestTable

// populateGet populates test cases for the Get function of the fileFault struct
func (tt *TestTable) populateGet() {
	testCases := make(map[string]string)
	testCases["test:1"] = "test:1"
	testCases["test:1,test2:2,test3:3"] = "test2:2"
	testCases["test:1,test2:2,test3:3,test2:a"] = "test2:a"
	testCases["test:1,test2:2,test3:3"] = "test4:"
	testCases["1:1test,test2:2,3:3"] = "1:1test"
	testCases["1:1test,&:2,2:&"] = "&:2"

	tt.get = testCases
}

// populateGetAll populates test cases for the GetAll function of the fileFault struct
func (tt *TestTable) populateGetAll() {
	testCases := make(map[string]string)
	testCases["test:1"] = "test:1"
	testCases["test:1,test2:2,test3:3"] = "test:1,test2:2,test3:3"
	testCases["test:1,test2:2,test3:3,test2:a"] = "test:1,test3:3,test2:a"
	testCases["1:a,d:2,test1:3,&(a0:a"] = "1:a,d:2,test1:3,&(a0:a"

	tt.getAll = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateGet()
	tt.populateGetAll()
	return tt
}

const testEncodingKey = "6368616e676520746869732070617373"
const testFilePath = "./test"

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
