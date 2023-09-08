package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler serves the form.html file for GET requests and processes the form data for POST requests
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Serve the static form.html file
		http.ServeFile(w, r, "./static/form.html")
		return
	} else if r.Method == "POST" {
		// Parse form data from the request
		if err := r.ParseForm(); err != nil {
			// Send error if there's an issue parsing
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// Display submitted form data
		fmt.Fprintf(w, "Post from website!\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	}
}

// helloHandler serves a simple "Hello!" response for GET requests to /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		// Return 404 if the URL path is not /hello
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		// Allow only GET method for this handler
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Send "Hello!" response
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// Create a new file server for serving static files
	fileServer := http.FileServer(http.Dir("./static"))
	// Handle all requests using the file server
	http.Handle("/", fileServer)
	// Register the formHandler for the path /form
	http.HandleFunc("/form", formHandler)
	// Register the helloHandler for the path /hello
	http.HandleFunc("/hello", helloHandler)

	// Log the start of the server
	fmt.Println("Starting server at port 8080")
	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If there's an error starting the server, log it and exit
		log.Fatal(err)
	}
}
