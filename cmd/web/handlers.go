package main

import (
		"fmt"
		"net/http"
		"strconv"
)

//handler definitions
func getHome(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "net/http Golang") //for adding additional header details
		w.Header().Add("Author", "nimilgp") 
		w.Header().Add("Program", "Snippet-Box") 
		w.Header().Add("Field-Name", "Can't have spaces in them") 
		fmt.Fprintf(w, "Hello from snippet-box!!!")
}


func getSnippetView(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id<1 {
				http.NotFound(w,r)
				return
		}
		fmt.Fprintf(w, "Display Snippet no: %d", id)
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Create snippet")
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) //201 created status code
		fmt.Fprintf(w, "POST part of create snippet")
}

