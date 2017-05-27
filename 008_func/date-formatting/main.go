package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}

}
