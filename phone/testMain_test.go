package phone

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

type TestTable struct {
	complexPhone       map[string]string
	normalizePhoneData map[string]string
}

// populateComplexPhone populates test cases for complex data transaction simulations of the taskDatabase struct.
func (tt *TestTable) populateComplexPhone() {
	testCases := make(map[string]string)
	testCases["#i,1,x"] = "x"
	testCases["#i,3,x,y,z,#u,1,2,yy,#d,1,1"] = "yy,z"
	testCases["#i,5,v,w,x,y,z,#u,3,1,a,1,b,2,c,#d,2,3,4"] = "b,c,z"
	testCases["#i,3,v,w,x,#d,1,1,#u,1,2,a"] = "a,x"
	testCases["#i,2,v,w,#d,1,1,#i,2,x,y,#u,2,2,a,3,c,#d,1,4"] = "a,c"

	tt.complexPhone = testCases
}

// populateNormalizePhoneData populates test cases for the normalizePhoneData function of the taskDatabase struct.
func (tt *TestTable) populateNormalizePhoneData() {
	testCases := make(map[string]string)
	testCases["#i,1,1234567890"] = "1234567890"
	testCases["#i,3,123-456-7890,(123)4567891,123-4567891"] = "1234567890,1234567891"
	testCases["#i,3,123 456 7890,(123) 456 7892,(123) 456-7893"] = "1234567890,1234567892,1234567893"
	testCases["#i,1,_1 :2#$(@34av56fF!7}8Q90"] = "1234567890"

	tt.normalizePhoneData = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateComplexPhone()
	tt.populateNormalizePhoneData()
	return tt
}

var testTable TestTable
var testPhoneDb *phoneDatabase

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	db, err := sql.Open("postgres", "user=postgres dbname=Custom sslmode=disable password=")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	testPhoneDb, err = NewPhoneDatabase(db)
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
