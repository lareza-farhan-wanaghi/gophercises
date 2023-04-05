package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	Data map[string]interface{}
	API  string
}

//go:embed templates
var templateFS embed.FS

// renderTemplate renders a template for a given page
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)
	t, ok := app.templateCache[templateToRender]
	if !ok {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			return err
		}
	}

	err = t.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}

// parseTemplate creates a template based on a given page name, base layout, and partial names
func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}
	if err != nil {
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
