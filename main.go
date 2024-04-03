package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Write the response header with content type as HTML
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML response
	fmt.Fprintf(w, "<html><body><h1>Hello, World!</h1></body></html>")
}

func main() {
	// Register a handler function for the root URL path "/"
	http.HandleFunc("/", helloWorldHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
