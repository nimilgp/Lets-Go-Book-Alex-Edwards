package main

import "net/http"

func (app *application) routes(staticDir string) *http.ServeMux {
		mux := http.NewServeMux()

		//create file server
		fileServer := http.FileServer(http.Dir(staticDir))
		
		//register mux.Handle func to register yje file server asd the handler for url path /static
		mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

		//register url patterns with handlers
		mux.HandleFunc("GET /{$}", app.getHome)
		mux.HandleFunc("GET /snippet/view/{id}", app.getSnippetView)
		mux.HandleFunc("GET /snippet/create", app.getSnippetCreate)
		mux.HandleFunc("POST /snippet/create", app.postSnippetCreate)

		return mux
}
