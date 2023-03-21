package cyoa

import (
	"net/http"
)

// Default acts as a generic handler that handles every route
func (app *application) Default(w http.ResponseWriter, r *http.Request) {
	arc, ok := app.cache["arcmap"].(map[string]Arc)[r.URL.Path[1:]]
	if !ok {
		http.Redirect(w, r, "intro", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["arc-title"] = arc.Title
	data["arc-story"] = arc.Story
	data["arc-options"] = arc.Options
	err := app.renderTemplate(w, r, "home", &templateData{
		Data: data,
	})
	if err != nil {
		panic((err))
	}
}
