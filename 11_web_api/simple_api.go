package main

import (
	"fmt"
	"net/http"
	"log"
)

// This example shows how to build a basic web server using Go's standard library.

// A handler function for the root URL `/`.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// `w` is the response writer, where you write your HTTP response.
	// `r` is the request, which contains details about the incoming request.

	fmt.Fprintf(w, "Welcome to the Simple Go API!")
}

// A handler function for the `/hello` endpoint.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// The URL query string can be accessed via `r.URL.Query()`.
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// A handler for a JSON response.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// You should set the content type header for JSON responses.
	w.Header().Set("Content-Type", "application/json")
	jsonResponse := `{"message": "This is a JSON response"}`
	fmt.Fprint(w, jsonResponse)
}

func main() {
	// --- Building a Simple Web Server ---
	// To run this, open your terminal, navigate to this directory, and run `go run simple_api.go`.
	// Then, open your browser and go to http://localhost:8080 or use a tool like `curl`.

	// `http.HandleFunc` registers a handler function for a given URL path.
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/json", jsonHandler)

	// `http.ListenAndServe` starts the HTTP server. It listens on a network address.
	// The second argument is the handler, which is `nil` here, meaning it uses the default `http.ServeMux`.
	// The `log.Fatal` call will cause the program to exit if `ListenAndServe` returns an error (e.g., if the port is already in use).
	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
