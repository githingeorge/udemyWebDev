package main

import (
	"html/template"
	"net/http"

	"github.com/githingeorge/udemyWebDev/034_photo-blog/cookie/session"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := session.GetCookie(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", cookie.Value)

}
