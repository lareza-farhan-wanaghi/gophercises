package pdf

import (
	"errors"
	"os"
	"testing"
	"time"
)

// TestCertCreate tests the Create function of the certPDF struct
func TestCertCreate(t *testing.T) {
	for k, v := range testTable.certCreate {
		certPDF := NewCertPDF()
		err := certPDF.Create(k, "test", time.Now())
		if v == "error" {
			if err == nil {
				t.Fatalf("expected an error but got nil. k:%s v:%s", k, v)
			}
			continue
		}

		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected a file with a path of %s. k:%s v:%s", v, k, v)
		}

		err = os.RemoveAll(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// TestInvoiceCreate tests the Create function of the invoicePDF struct
func TestInvoiceCreate(t *testing.T) {
	for k, v := range testTable.invoiceCreate {
		invoicePDF := NewInvoicePDF()

		mockItems := []*InvoiceItem{
			{Name: "Test1", PricePerUnit: 10, Quantity: 10},
			{Name: "Test2", PricePerUnit: 20, Quantity: 20},
			{Name: "Test3", PricePerUnit: 30, Quantity: 30},
		}
		err := invoicePDF.Create(k, mockItems)
		if v == "error" {
			if err == nil {
				t.Fatalf("expected an error but got nil. k:%s v:%s", k, v)
			}
			continue
		}

		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(v); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("expected a file with a path of %s. k:%s v:%s", v, k, v)
		}

		err = os.RemoveAll(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}
