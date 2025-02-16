package main

import (
	"fmt"
	"net/http"
)

var counter int

// Ignore favicon requests
func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "You are visitor number %d\n", counter)

	if len(r.URL.Path) > 1 {
		subPath := r.URL.Path[1:2]
		fmt.Fprintf(w, "First character of path is: %s\n", subPath)
	} else {
		fmt.Fprintf(w, "Path is too short to extract first character\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}
