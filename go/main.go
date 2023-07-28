package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Add your desired HTML content
	fmt.Fprintf(w, "<!DOCTYPE html>")
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head>")
	fmt.Fprintf(w, "<title>My Go Webpage</title>")
	fmt.Fprintf(w, "<style>")
	// Add your CSS code here
	fmt.Fprintf(w, "body { font-family: Arial, sans-serif; }")
	fmt.Fprintf(w, "h1 { color: #007bff; }")
	fmt.Fprintf(w, "p { color: #333; }")
	fmt.Fprintf(w, "</style>")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<h1>Hello World!, This Go application runs on Triggr</h1>")
	fmt.Fprintf(w, "<p>Welcome to my webpage.</p>")
	fmt.Fprintf(w, "<p>This is a sample Go application's webpage.</p>")
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 3007...")
	http.ListenAndServe(":3007", nil)
}
