package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("Hello from SnippetBox"))
	if err != nil {
		return
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a particular snippet"))
	if err != nil {
		return
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	_, err := w.Write([]byte("Create a new snippet"))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
