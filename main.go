package main

import (
  "fmt"
  "strconv"
  "log"
  "net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>Hello!, from the webserver</h1>"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil || id < 1 {
    http.NotFound(w,r)
    return
  } 

  msg := fmt.Sprintf("<h1>View snippet number:%d</h1>", id)
  w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>Create a snippet</h1>"))
}

func main(){
  mux := http.NewServeMux()

  mux.HandleFunc("/{$}", root)
  mux.HandleFunc("/snippet/view/{id}", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  log.Println("Hello, world!")

  if err := http.ListenAndServe(":3333", mux); err != nil {
    log.Fatal(err)
  }
}
