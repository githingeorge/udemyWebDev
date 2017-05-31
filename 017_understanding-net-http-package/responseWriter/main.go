package main

import "net/http"
import "fmt"

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Githin-Key", "This is from Githin")
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}
func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}
