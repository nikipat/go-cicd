package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc Sample documentation for HandlerFunc
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nRelease 1")
	fmt.Fprintf(w, "\nRelease 2")
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
