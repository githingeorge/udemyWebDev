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

func main() {
	strings := []string{"hi", "githin", "george", "buddha", "gandhi"}
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", strings)
	if err != nil {
		log.Fatal(err)
	}

}
