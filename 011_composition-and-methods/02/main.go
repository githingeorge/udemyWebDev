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

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := person{
		"James Bond",
		42,
	}
	agent := doubleZero{
		p1,
		true,
	}
	err := tmpl.Execute(os.Stdout, agent)
	if err != nil {
		log.Fatal(err)
	}
}
