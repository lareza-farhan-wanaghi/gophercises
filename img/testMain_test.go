package img

import (
	"os"
	"testing"
)

type TestTable struct {
	calculatedWidth  map[string]string
	calculatedHeight map[string]string
	drawSvg          map[string]string
	drawPng          map[string]string
}

var testTable TestTable

// populateDrawSvg populates test cases for the DrawSvg function of the chart struct
func (tt *TestTable) populateDrawSvg() {
	testCases := map[string]string{
		"test.png.svg":  "test.png.svg",
		"test1.svg":     "test1.svg",
		"test.png":      "error",
		"test1.svg.png": "error",
	}

	tt.drawSvg = testCases
}

// populateDrawPng populates test cases for the DrawPng function of the chart struct
func (tt *TestTable) populateDrawPng() {
	testCases := map[string]string{
		"test.png.svg":  "error",
		"test1.svg":     "error",
		"test.png":      "test.png",
		"test1.svg.png": "test1.svg.png",
	}

	tt.drawPng = testCases
}

// populateCalculatedWidth populates test cases for the calculated width of the chart struct
func (tt *TestTable) populateCalculatedWidth() {
	testCases := map[string]string{
		"1,10,5":    "20",
		"2,15,5":    "45",
		"1,15,10":   "35",
		"4,15,10":   "110",
		"3,100,500": "2300",
	}

	tt.calculatedWidth = testCases
}

// populateCalculatedHeight populates test cases for the calculated height of the chart struct
func (tt *TestTable) populateCalculatedHeight() {
	testCases := map[string]string{
		"1,2,3,10,10":       "40",
		"1,3,5,15,15":       "90",
		"1,1,1,1,1,25,1,10": "35",
		"7,1,3,5,2,1,10,5":  "75",
	}

	tt.calculatedHeight = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateCalculatedWidth()
	tt.populateCalculatedHeight()
	tt.populateDrawSvg()
	tt.populateDrawPng()

	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
