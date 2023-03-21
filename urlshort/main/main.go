package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/lareza-farhan-wanaghi/gophercises/urlshort"
	_ "github.com/lib/pq"
)

// main provides the entry point of the app
func main() {
	mux := defaultMux()
	databaseFlag := flag.String("d", "", "Specifies the database address that will be used. If isn't provided, will use file instead")
	fileFlag := flag.String("f", "pathyaml.yaml", "Specifies the file containing mapped paths")
	flag.Parse()

	var handler http.HandlerFunc
	var err error
	if len(*databaseFlag) == 0 {
		handler, err = fileBasedHandler(fileFlag, mux)
		if err != nil {
			panic(err)
		}
	} else {
		db, err := sql.Open("postgres", *databaseFlag)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		handler, err = databaseBasedHandler(mux, db)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

// fileBasedHandler creates a handler based on a file
func fileBasedHandler(fileFlag *string, fallback http.Handler) (http.HandlerFunc, error) {
	var handler http.HandlerFunc
	if filepath.Ext(*fileFlag) == ".yaml" {
		yaml, err := ioutil.ReadFile(*fileFlag)
		if err != nil {
			return nil, err
		}
		handler, err = urlshort.YAMLHandler(yaml, fallback)
		if err != nil {
			return nil, err
		}
	} else if filepath.Ext(*fileFlag) == ".json" {
		json, err := ioutil.ReadFile(*fileFlag)
		if err != nil {
			return nil, err
		}
		handler, err = urlshort.JSONHandler(json, fallback)
		if err != nil {
			return nil, err
		}
	} else {
		panic("The provided file format is currently not supported.")
	}
	return handler, nil
}

// databaseBasedHandler creates a handler based on a database
func databaseBasedHandler(fallback http.Handler, db *sql.DB) (http.HandlerFunc, error) {
	handler, err := urlshort.DatabaseHandler(fallback, db)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

// defaultMux creates a default handler
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}
