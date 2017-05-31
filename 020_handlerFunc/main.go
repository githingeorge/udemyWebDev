package main

import "net/http"
import "io"

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "meow meow")
}
func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandleFunc(c))
	http.HandleFunc("*", func(w http.ResponseWriter, res *http.Request) {
		io.WriteString(w, "404")

	})
	http.ListenAndServe(":8000", nil)
}
