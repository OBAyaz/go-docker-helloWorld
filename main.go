package main

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":3333", nil)
}
