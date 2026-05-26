package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from snippetbox"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createPostSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Display a form for creating a new snippet..."))
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", viewSnippet)
	mux.HandleFunc("GET /snippet/create", createSnippet)
	mux.HandleFunc("POST /snippet/create", createPostSnippet)
	log.Println("Starting server at :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
