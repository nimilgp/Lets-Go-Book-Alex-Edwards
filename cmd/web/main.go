package main

import (
		"log"
		"net/http"
)

func main() {
		port := ":3333"

		//create a new serveMux
		mux := http.NewServeMux()

		//create file server
		fileServer := http.FileServer(http.Dir("./ui/static/"))
		
		//register mux.Handle func to register yje file server asd the handler for url path /static
		mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

		//register url patterns with handlers
		mux.HandleFunc("GET /{$}", getHome)
		mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
		mux.HandleFunc("GET /snippet/create", getSnippetCreate)
		mux.HandleFunc("POST /snippet/create", postSnippetCreate)

		log.Println("Server is up on port ",port)
		
		//start a new web server at a port, handled by a serveMux
		err := http.ListenAndServe(port, mux)
		log.Fatal(err)
}
