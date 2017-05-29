package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
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

	ford := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}

	toyota := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi}
	cars := []car{ford, toyota}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
