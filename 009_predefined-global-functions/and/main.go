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

type sage struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: true,
	}

	gandhi := sage{
		Name:  "gandhi",
		Motto: "Be the change",
		Admin: false,
	}

	user := sage{
		Name:  "",
		Motto: "Be the change",
		Admin: true,
	}

	sages := []sage{buddha, gandhi, user}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}

}
