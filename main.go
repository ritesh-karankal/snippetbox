package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Display a form for creating a new snippet..."))
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server at :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
