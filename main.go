package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "index")
	} else if r.URL.Path == "/about" {
		fmt.Fprintln(w, "/about")
	} else {
		fmt.Fprintln(w, "error")
	}
}

func main() {
	http.HandleFunc("/", HandleFunc)
	http.ListenAndServe(":3000", nil)
}
