package main

import (
	"time"

	"github.com/lareza-farhan-wanaghi/gophercises/pdf"
)

// main provides the entry point of the app
func main() {
	invociePDF := pdf.NewInvoicePDF()
	err := invociePDF.Create("demoInvoice.pdf", []*pdf.InvoiceItem{
		{Name: "Test", PricePerUnit: 20, Quantity: 3},
		{Name: "Test Space", PricePerUnit: 120, Quantity: 7},
		{Name: "Test Long String Test Long String Test Long String", PricePerUnit: 10, Quantity: 2},
	})
	if err != nil {
		panic(err)
	}

	certPDF := pdf.NewCertPDF()
	err = certPDF.Create("demoCert.pdf", "Lareza Farhan Wanaghi", time.Now())
	if err != nil {
		panic(err)
	}
}
