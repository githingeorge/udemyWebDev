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

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Summer, Spring semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps Using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
			},
		},
	}
	err := tmpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}
}
