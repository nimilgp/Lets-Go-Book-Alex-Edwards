package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Application", "Paste Bin")

	ts, err := template.ParseFiles("./ui/html/pages/root.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "<h1>View snippet number:%d</h1>", id)
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Display a form to create a snippet</h1>")
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "<h1>Save a new snippet</h1>")
}
