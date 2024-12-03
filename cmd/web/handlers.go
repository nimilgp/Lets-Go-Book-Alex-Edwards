package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/nimilgp/paste-bin/internal/models"
)

func (app *application) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Application", "Paste Bin")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "root.tmpl.html", templateData{
		Snippets: snippets,
	})
}

func (app *application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.tmpl.html", templateData{
		Snippet: snippet,
	})
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Display a form to create a snippet</h1>")
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
