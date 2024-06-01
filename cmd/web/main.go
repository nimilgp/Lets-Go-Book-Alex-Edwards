package main

import (
		"log"
		"net/http"
)

func main() {
		//create a new serveMux
		mux := http.NewServeMux()

		//register url patterns with handlers
		mux.HandleFunc("GET /{$}", getHome)
		mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
		mux.HandleFunc("GET /snippet/create", getSnippetCreate)
		mux.HandleFunc("POST /snippet/create", postSnippetCreate)

		log.Println("Server is up!")
		
		//start a new web server at a port, handled by a serveMux
		err := http.ListenAndServe(":3333", mux)
		log.Fatal(err)
}
