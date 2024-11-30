package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := ":3333"
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Println("server running at port", PORT)

	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatal(err)
	}
}
