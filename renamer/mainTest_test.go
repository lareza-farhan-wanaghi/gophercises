package renamer

import (
	"os"
	"testing"
)

type TestTable struct {
	renameFile map[string]string
}

var testTable TestTable

// populateRenameFile populates test cases for the renameFile function of the gameManager
func (tt *TestTable) populateRenameFile() {
	testCases := map[string]string{}
	testCases["a_a,b_b,c_c,d/d_d,_,"] = "aa,bb,cc,dd"
	testCases["a_a,b_b,c_c,d/d_d,_,a"] = "aaa,bab,cac,dad"
	testCases["a_01,b_02,c_03,[0-9],a"] = "a_aa,b_aa,c_aa"
	testCases["a_01,b_02,c_03,[0-9]+,"] = "a_,b_,c_"
	testCases["a_01,b_02,c_03,[0-9]+,aaaaa"] = "a_aaaaa,b_aaaaa,c_aaaaa"
	testCases["a/a/a_a,b/b_b,c_c,a/d_d,_,blob"] = "abloba,dblobd,bblobb,cblobc"
	testCases["aa,a_a,_,"] = "a_a,aa"

	tt.renameFile = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateRenameFile()

	return tt
}

var testDir string

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()
	testDir = "./test/"

	exitVal := m.Run()
	os.Exit(exitVal)
}
