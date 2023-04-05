package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/lareza-farhan-wanaghi/gophercises/urlshort"
	_ "github.com/lib/pq"
)

// main provides the entry point of the app
func main() {
	mux := urlshort.GetDefaultHandler()
	databaseFlag := flag.String("d", "", "Specifies the database address that will be used. If it isn't provided, will use file instead")
	fileFlag := flag.String("f", "pathyaml.yaml", "Specifies the file containing mapped paths")
	flag.Parse()

	var handler http.HandlerFunc
	var err error
	if len(*databaseFlag) == 0 {
		handler, err = urlshort.GetFileBasedHandler(fileFlag, mux)
		if err != nil {
			panic(err)
		}
	} else {
		db, err := sql.Open("postgres", *databaseFlag)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		handler, err = urlshort.GetDatabaseBasedHandler(mux, db)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}
