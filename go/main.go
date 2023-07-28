package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Add your desired HTML content
	htmlContent := `
<!DOCTYPE html>
<html>
<head>
  <title>My Go Webpage</title>
  <style>
    /* Add your CSS code here */
    body {
      font-family: Arial, sans-serif;
      background-color: #e6f7ff;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
      margin: 0;
    }

    h1 {
      color: #007bff;
    }

    p {
      color: #333;
    }
  </style>
</head>
<body>
  <div>
    <h1>Hello World!, This Go application runs on Triggr</h1>
    <p>Welcome to my webpage.</p>
    <p>This is a sample Go application's webpage.</p>
  </div>
</body>
</html>
`

	// Write the HTML content to the response
	fmt.Fprintf(w, htmlContent)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 3007...")
	http.ListenAndServe(":3007", nil)
}
