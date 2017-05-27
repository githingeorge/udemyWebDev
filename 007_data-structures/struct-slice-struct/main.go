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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
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

	// data := items{
	// 	Wisdom:    sages,
	// 	Transport: cars,
	// }

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
