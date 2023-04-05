package task

import (
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
)

type TestTable struct {
	addTask     map[string]string
	doTask      map[string]string
	deleteTask  map[string]string
	complexTask map[string]string
}

var testTable TestTable

// populateAddTask populates test cases for the addTask function of the taskDatabase struct
func (tt *TestTable) populateAddTask() {
	testCases := make(map[string]string)
	testCases["cleaning rooms"] = "cleaning rooms"
	testCases["cleaning rooms,cooking eggs"] = "cleaning rooms,cooking eggs"
	testCases["cleaning rooms,cleaning rooms,cleaning rooms"] = "cleaning rooms"
	testCases["Buying 1 egg,Sending $30 to mom"] = "Buying 1 egg,Sending $30 to mom"

	tt.addTask = testCases
}

// populateDoTask populates test cases for the doTask function of the taskDatabase struct
func (tt *TestTable) populateDoTask() {
	testCases := make(map[string]string)
	testCases["1,cleaning rooms,0"] = "cleaning rooms"
	testCases["3,x,y,z,1"] = "y"
	testCases["1,x,1"] = ""
	testCases["1,x,-1"] = ""

	tt.doTask = testCases
}

// populateDoTask populates test cases for the doTask function of the taskDatabase struct
func (tt *TestTable) populateDeleteTask() {
	testCases := make(map[string]string)
	testCases["1,cleaning rooms,0"] = "cleaning rooms"
	testCases["3,x,y,z,1"] = "y"
	testCases["1,x,1"] = ""
	testCases["1,x,-1"] = ""

	tt.deleteTask = testCases
}

// populateDoTask populates test cases for the doTask function of the taskDatabase struct.
func (tt *TestTable) populateComplexTask() {
	testCases := make(map[string]string)
	testCases["1,cleaning rooms,1,0,0"] = "1,cleaning rooms,0"
	testCases["3,x,y,z,1,1,0"] = "1,y,2,x,z"
	testCases["1,x,0,1,0"] = "0,0"
	testCases["5,v,w,x,y,z,2,2,2,1,1"] = "2,x,y,2,v,z"
	testCases["5,v,w,x,y,z,2,1,2,0"] = "2,w,y,3,v,x,z"
	testCases["5,v,w,x,y,z,2,-1,2,0"] = "1,x,4,v,w,y,z"

	tt.complexTask = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateAddTask()
	tt.populateDoTask()
	tt.populateComplexTask()
	tt.populateDeleteTask()
	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	dbName := "mock.db"
	err := os.Remove(dbName)
	if err != nil {
		panic(err)
	}

	dbOps := &bolt.Options{Timeout: 1 * time.Second}
	db, err := bolt.Open(dbName, 0600, dbOps)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskDb, err = newTaskDatabase(db)
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
