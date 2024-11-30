package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", "3333", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Println("server running at port:", cfg.addr)

	if err := http.ListenAndServe(":"+cfg.addr, mux); err != nil {
		log.Fatal(err)
	}
}
