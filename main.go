package main

import (
		"fmt"
		"log"
		"net/http"
		"strconv"
)

//handler definitions
func getHome(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "net/http Golang") //for adding additional header details
		w.Header().Add("Author", "nimilgp") 
		w.Header().Add("Program", "Snippet-Box") 
		w.Header().Add("Field-Name", "Can't have spaces in them") 
		w.Write([]byte("Hello from snippet-box!!!"))
}


func getSnippetView(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id<1 {
				http.NotFound(w,r)
				return
		}
		msg := fmt.Sprintf("Display snippet no: %d", id)
		w.Write([]byte(msg))
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create snippet"))
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) //201 created status code
		w.Write([]byte("POST part of create snippet"))
}

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
