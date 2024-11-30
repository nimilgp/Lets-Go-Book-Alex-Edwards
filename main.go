package main

import (
  "fmt"
  "strconv"
  "log"
  "net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Server", "Go")
  w.Header().Add("Application", "Snippet Box")
  fmt.Fprintf(w, "<h1>Hello!, from the webserver</h1>")
}

func getSnippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil || id < 1 {
    http.NotFound(w,r)
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

func main(){
  PORT := ":3333"
  mux := http.NewServeMux()

  mux.HandleFunc("GET /{$}", getRoot)
  mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
  mux.HandleFunc("GET /snippet/create", getSnippetCreate)
  mux.HandleFunc("POST /snippet/create", postSnippetCreate)

  log.Println("server running at port",PORT)

  if err := http.ListenAndServe(PORT, mux); err != nil {
    log.Fatal(err)
  }
}
