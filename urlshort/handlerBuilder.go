package urlshort

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// getFileBasedHandler creates a handler based on a file
func GetFileBasedHandler(fileFlag *string, fallback http.Handler) (http.HandlerFunc, error) {
	var handler http.HandlerFunc
	if filepath.Ext(*fileFlag) == ".yaml" {
		yaml, err := ioutil.ReadFile(*fileFlag)
		if err != nil {
			return nil, err
		}
		handler, err = YAMLHandler(yaml, fallback)
		if err != nil {
			return nil, err
		}
	} else if filepath.Ext(*fileFlag) == ".json" {
		json, err := ioutil.ReadFile(*fileFlag)
		if err != nil {
			return nil, err
		}
		handler, err = JSONHandler(json, fallback)
		if err != nil {
			return nil, err
		}
	} else {
		panic("The provided file format is currently not supported.")
	}
	return handler, nil
}

// databaseBasedHandler creates a handler based on a database
func GetDatabaseBasedHandler(fallback http.Handler, db *sql.DB) (http.HandlerFunc, error) {
	handler, err := DatabaseHandler(fallback, db)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

// GetDefaultHandler creates a default handler
func GetDefaultHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}
