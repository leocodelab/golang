package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":3000", nil)
}
