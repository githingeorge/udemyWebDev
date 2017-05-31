package main

import (
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog.jpg", dogPic)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "foo ran")
}

func dog(rw http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("dog.gohtml"))
	tmpl.ExecuteTemplate(rw, "dog.gohtml", nil)
}

func dogPic(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "dog.jpg")

}
