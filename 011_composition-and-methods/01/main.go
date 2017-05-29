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

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		"James Bond",
		42,
	}
	err := tmpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
