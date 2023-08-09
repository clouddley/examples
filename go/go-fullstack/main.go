package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux
	r := http.NewServeMux()

	// Serve static files from the "static" directory with "/static/" prefix
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up the root handler to serve the "index.html" file
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Set up the HTTP server
	server := &http.Server{
		Addr:    ":3007",
		Handler: r,
	}

	// Start the server
	fmt.Println("Server is listening on port 3007...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
