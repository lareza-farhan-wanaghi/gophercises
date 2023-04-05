package main

import "html/template"

// main provides the entry point of the app
func main() {
	app := &application{
		templateCache: make(map[string]*template.Template),
		cache:         make(map[string]interface{}),
		api:           "http://localhost:8080",
	}

	err := app.serve()
	if err != nil {
		panic(err)
	}
}
