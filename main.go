package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc1 Sample documentation for HandlerFunc
func HandlerFunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nRelease 0")
}

// //Calculate just calculates
// func Calculate(val int)(result int){
// 	return val*2
// }

func main() {
	http.HandleFunc("/", HandlerFunc1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
