package main

import (
  "log"
  "net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>hello from webserver!</h1>"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>View specific snippet</h1>"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>Create a snippet</h1>"))
}

func main(){
  mux := http.NewServeMux()

  mux.HandleFunc("/", root)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  log.Println("Hello, world!")

  if err := http.ListenAndServe(":3333", mux); err != nil {
    log.Fatal(err)
  }
}
