package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "base.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}
