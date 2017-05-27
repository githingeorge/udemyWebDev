package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "gandhi",
		Motto: "Be the change",
	}

	err := tpl.Execute(os.Stdout, []sage{buddha, gandhi})
	if err != nil {
		log.Fatal(err)
	}
}
