package main

import (
		"flag"
		"log"
		"net/http"
)

func main() {
		portptr := flag.String("port", ":3333", "HTTP port address")//flag, default val, helper
		flag.Parse()//call this before use of the flag variables else will stay at default

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

		log.Println("Server Port given is ", *portptr)
		
		//start a new web server at a port, handled by a serveMux
		err := http.ListenAndServe(*portptr, mux)
		log.Fatal(err)
}
