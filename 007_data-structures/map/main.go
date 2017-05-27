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
func main() {
	sages := map[string]string{
		"India":    "Gabdhi",
		"America":  "MLK",
		"Meditate": "Buddha",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}
}
