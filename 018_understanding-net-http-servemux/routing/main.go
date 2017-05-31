package main

import "net/http"
import "io"

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy")
	case "/cat":
		io.WriteString(res, "meow meow")

	default:
		io.WriteString(res, "404")
	}
}
func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}
