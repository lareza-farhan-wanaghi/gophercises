package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lareza-farhan-wanaghi/gophercises/transform"
)

// HomeHandler responds to the request by rendering the home page
func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, r, "home", &templateData{})
	if err != nil {
		app.writeError(w, err, http.StatusInternalServerError)
		return
	}
}

// SampleModesHandler responds to the request by returning JSON data containing primitive-shape-brushed images with different mode variables
func (app *application) SampleModesHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	tmpDir := "images/tmp"
	err = os.MkdirAll(tmpDir, 0777)
	if err != nil {
		app.writeError(w, err, http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir)

	filePath, err := app.saveAttachedImage(w, r, "file", tmpDir)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	payloadData := []transformData{}
	modes := [9]string{"Combo", "Triangle", "Rect", "Ellipse", "Circle", "Rotated Rect", "Beziers", "Rotated Ellipse", "Polygon"}
	n := 100
	extIndex := strings.LastIndex(filePath, ".")
	for i, mode := range modes {
		out := fmt.Sprintf("%s-out-%d-%d%s", filePath[:extIndex], i, n, filePath[extIndex:])
		err := transform.Primitive(filePath, out, i, n)
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}

		image, err := app.readFile(out)
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}

		payloadData = append(payloadData, transformData{Name: mode, Image: image})
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}
	}

	app.writeJSON(w, http.StatusOK, payload{Error: false, Data: payloadData})
}

// SampleNsHandler responds to the request by returning JSON data containing primitive-shape-brushed images with different n variables
func (app *application) SampleNsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	tmpDir := "images/tmp"
	err = os.MkdirAll(tmpDir, 0777)
	if err != nil {
		app.writeError(w, err, http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir)

	filePath, err := app.saveAttachedImage(w, r, "file", tmpDir)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	mQuesry := r.URL.Query().Get("m")
	modeIndex, err := strconv.Atoi(mQuesry)
	if err != nil {
		app.writeError(w, err, http.StatusBadRequest)
		return
	}

	payloadData := []transformData{}
	extIndex := strings.LastIndex(filePath, ".")
	nStart := 25
	nInterval := 25
	for i := 0; i < 6; i++ {
		n := nStart + nInterval*i
		out := fmt.Sprintf("%s-out-%d-%d%s", filePath[:extIndex], modeIndex, n, filePath[extIndex:])
		err := transform.Primitive(filePath, out, modeIndex, n)
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}

		image, err := app.readFile(out)
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}

		payloadData = append(payloadData, transformData{Name: fmt.Sprintf("%dn", n), Image: image})
		if err != nil {
			app.writeError(w, err, http.StatusInternalServerError)
			return
		}
	}

	app.writeJSON(w, http.StatusOK, payload{Error: false, Data: payloadData})
}
