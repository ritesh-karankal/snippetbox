package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
        w.Write([]byte(msg))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Display a form for creating a new snippet..."))
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server at :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
