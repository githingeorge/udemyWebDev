package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

type Page struct {
	Title, Heading, Input string
}

func main() {
	home := Page{
		Title:   "Nothing is escaped",
		Heading: "Nothing is escaped",
		Input:   "<script>alert('yo alert')</script>",
	}
	err := tmpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatal(err)
	}
}
