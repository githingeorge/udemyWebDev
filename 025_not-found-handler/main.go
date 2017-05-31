package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(rw, "go look at ur termial")
}
