package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/lareza-farhan-wanaghi/gophercises/cyoa"
)

type application struct {
	templateCache map[string]*template.Template
	cache         map[string]interface{}
}

// serve serves a http listener
func (app *application) serve() error {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	fmt.Printf("Starting HTTP server on port %s\n", srv.Addr)
	return srv.ListenAndServe()
}

// main provides the entry point of the app
func main() {
	app := &application{
		templateCache: make(map[string]*template.Template),
		cache:         make(map[string]interface{}),
	}

	fileFlag := flag.String("f", "gopher.json", "Specifies the file of the story tree")
	flag.Parse()
	arcMap, err := cyoa.GetArcMap(*fileFlag)
	if err != nil {
		panic(err)
	}
	app.cache["arcmap"] = arcMap

	err = app.serve()
	if err != nil {
		panic(err)
	}
}
