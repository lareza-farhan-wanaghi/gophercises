package pdf

import (
	"os"
	"testing"
)

type TestTable struct {
	certCreate    map[string]string
	invoiceCreate map[string]string
}

var testTable TestTable

// populateCertCreate populates test cases for the Create function of the certPDF struct
func (tt *TestTable) populateCertCreate() {
	testCases := map[string]string{
		"test.pdf":       "test.pdf",
		"test1.pdf":      "test1.pdf",
		"2test.pdf":      "2test.pdf",
		"3test1.svg.pdf": "3test1.svg.pdf",
		"test4.pdf.png":  "error",
		"test4.png":      "error",
	}

	tt.certCreate = testCases
}

// populateInvoiceCreate populates test cases for the Create function of the invoicePDF struct
func (tt *TestTable) populateInvoiceCreate() {
	testCases := map[string]string{
		"test.pdf":       "test.pdf",
		"test1.pdf":      "test1.pdf",
		"2test.pdf":      "2test.pdf",
		"3test1.svg.pdf": "3test1.svg.pdf",
		"test4.pdf.png":  "error",
		"test4.png":      "error",
	}

	tt.invoiceCreate = testCases
}

// newTestTable creates a new TestTable
func newTestTable() TestTable {
	tt := TestTable{}
	tt.populateCertCreate()
	tt.populateInvoiceCreate()

	return tt
}

// TestMain provides the entry point of the test
func TestMain(m *testing.M) {
	testTable = newTestTable()

	exitVal := m.Run()
	os.Exit(exitVal)
}
