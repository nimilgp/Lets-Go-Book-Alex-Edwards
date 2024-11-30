package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Application", "Paste Bin")

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/root.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "<h1>View snippet number:%d</h1>", id)
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Display a form to create a snippet</h1>")
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "<h1>Save a new snippet</h1>")
}
