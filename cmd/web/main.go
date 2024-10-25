package main

import (
	"log"
	"net/http"
)

const PORT = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)

	log.Printf("Starting server on %s", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}
