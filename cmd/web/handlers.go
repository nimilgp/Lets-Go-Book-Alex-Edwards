package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Add("Application", "Snippet Box")
	fmt.Fprintf(w, "<h1>Hello!, from the webserver</h1>")
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
