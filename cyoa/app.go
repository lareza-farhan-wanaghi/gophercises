package cyoa

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
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

// Run runs the cyoa app web server
func Run() {
	app := &application{
		templateCache: make(map[string]*template.Template),
		cache:         make(map[string]interface{}),
	}

	arcMap, err := GetArcMap()
	if err != nil {
		panic(err)
	}
	app.cache["arcmap"] = arcMap

	err = app.serve()
	if err != nil {
		panic(err)
	}
}
