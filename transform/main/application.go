package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type application struct {
	templateCache map[string]*template.Template
	cache         map[string]interface{}
	api           string
}

type payload struct {
	Error bool `json:"error"`
	Data  any  `json:"data"`
}

// serve serves an http listener
func (app *application) serve() error {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           app.routes(),
		IdleTimeout:       45 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      0 * time.Second,
	}

	fmt.Printf("Starting HTTP server on port %s\n", srv.Addr)
	return srv.ListenAndServe()
}

// writeJSON writes aribtrary data out as JSON
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
		return
	}
}

// writeError writes a payload containing the error with the header modified
func (app *application) writeError(w http.ResponseWriter, err error, code int) {
	log.Printf("%v", err)
	payload := payload{
		Error: true,
		Data:  err.Error(),
	}

	app.writeJSON(w, code, payload)
}

// saveAttachedImage saves the image attached in the specified form fIle
func (app *application) saveAttachedImage(w http.ResponseWriter, r *http.Request, formFile string, outDir string) (string, error) {
	fIn, fHeader, err := r.FormFile(formFile)
	if err != nil {
		return "", err
	}
	defer fIn.Close()

	filePath := fmt.Sprintf("%s%s", time.Now().String(), fHeader.Filename)
	filePath = filepath.Join(outDir, filePath)
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, fIn)
	if err != nil {
		return "", err
	}

	log.Printf("saved an image on %s", filePath)
	return filePath, nil
}

// readFile reads a file specified by the path and returns the bytes read
func (app *application) readFile(path string) ([]byte, error) {
	out, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	var buff bytes.Buffer
	_, err = io.Copy(&buff, out)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
