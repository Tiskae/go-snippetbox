package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a spcific snippet..."))
}

// Add a snippetCreate handler function
func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Set a new cache-control header. If an existing "Cache-Control" header exists
		// it will be overwritten.		w.Header().Set("Allow", "POST")

		// Header manipulations

		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age-31536000")

		// Delete all values for the "Cache-Control" header.
		w.Header().Del("Cache-Control")

		// Retrieve the first value for the "Cache-Control" header.
		w.Header().Get("Cache-Control")

		// Retrieve a slice of all values for the "Cache-Control" header.
		w.Header().Values("Cache-Control")

		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// or...
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4040")
	err := http.ListenAndServe(":4040", mux)
	log.Fatal(err)
}
