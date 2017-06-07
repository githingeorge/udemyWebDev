package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", index)
	http.ListenAndServe(":8080", r)
}
func index(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("welcome"))
}
