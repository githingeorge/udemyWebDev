package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	tmpl.ExecuteTemplate(rw, "index.gohtml", nil)
}
