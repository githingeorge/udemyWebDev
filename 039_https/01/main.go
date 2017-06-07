package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("This is an example server. \n"))
}
