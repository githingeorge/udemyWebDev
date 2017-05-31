package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (hd hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Form)
	data := struct {
		Method      string
		Submissions url.Values
		Url         *url.URL
	}{
		Method:      req.Method,
		Submissions: req.Form,
		Url:         req.URL,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)

}
