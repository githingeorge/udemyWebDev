package main

import (
	"net/http"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("pic"))))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
