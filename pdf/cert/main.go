package main

import (
	"os"
	"time"

	"github.com/lareza-farhan-wanaghi/gophercises/pdf"
)

// main provides the entry point of the app
func main() {
	certPDF := pdf.NewCertPDF()
	err := certPDF.Create(os.Args[2], os.Args[1], time.Now())
	if err != nil {
		panic(err)
	}
}
