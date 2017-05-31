package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	fmt.Println(req.URL)
	err = tmpl.Execute(w, req.Form)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)

}
