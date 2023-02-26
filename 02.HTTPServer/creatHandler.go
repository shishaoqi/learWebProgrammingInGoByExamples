package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 01.HelloWorld
	http.HandleFunc("/", hello)

	// 02.HTTPServer
	h := new(Hello)
	http.Handle("/test", h)

	http.ListenAndServe(":80", nil)
}

type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " visit on: %s\n", r.URL.Path)
}

/*
	------------ 01.HelloWorld Content -------------

# The request handler

# A handler in Go is a function with this signature

An http.ResponseWriter which is where you write your text/html response to.
An http.Reqest which contains all information about this HTTP request including things like the URL or header fields.
*/
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

/* ------------ 02.HTTP Server Content -------------
引进的新内容有：
1. http.FileServer() 与 http.StripPrefix()
2. http.Handle()
*/
