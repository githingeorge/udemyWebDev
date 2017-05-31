package main

import "net/http"
import "io"

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy")
}

type coldcat int

func (m coldcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow meow")
}
func main() {
	var d hotdog
	var c coldcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	mux.HandleFunc("*", func(w http.ResponseWriter, res *http.Request) {
		io.WriteString(w, "404")

	})
	http.ListenAndServe(":8000", mux)
}
