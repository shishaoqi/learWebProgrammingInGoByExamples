package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":80", nil)
}

/*
The request handler

# A handler in Go is a function with this signature

An http.ResponseWriter which is where you write your text/html response to.
An http.Reqest which contains all information about this HTTP request including things like the URL or header fields.
*/
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
