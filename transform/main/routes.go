package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// routes specifies all routes within the app
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", app.HomeHandler)
	mux.Post("/api/sample-modes", app.SampleModesHandler)
	mux.Post("/api/sample-ns", app.SampleNsHandler)

	return mux
}
