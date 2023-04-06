package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/lareza-farhan-wanaghi/gophercises/pdf"
)

// main provides the entry point of the app
func main() {
	inpath := os.Args[1]
	outpath := os.Args[2]
	f, err := os.Open(inpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	invoiceItems := []*pdf.InvoiceItem{}
	for _, d := range data {
		invoicePPU, err := strconv.Atoi(d[1])
		if err != nil {
			panic(err)
		}

		invoiceQuantity, err := strconv.Atoi(d[2])
		if err != nil {
			panic(err)
		}

		invoiceItem := &pdf.InvoiceItem{
			Name:         d[0],
			PricePerUnit: invoicePPU,
			Quantity:     invoiceQuantity,
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}

	invociePDF := pdf.NewInvoicePDF()
	err = invociePDF.Create(outpath, invoiceItems)
	if err != nil {
		panic(err)
	}
}
