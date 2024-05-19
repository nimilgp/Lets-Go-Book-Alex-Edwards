package main

import (
		"log"
		"net/http"
)

//handler definitions
func home(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from snippet-box!!!"))
}

func main() {
		//create a new serveMux
		mux := http.NewServeMux()

		//register url patterns with handlers
		mux.HandleFunc("/", home)

		log.Println("Server is up!")
		//start a new web server at a port, handled by a serveMux
		err := http.ListenAndServe(":3333", mux)
		log.Fatal(err)
}
