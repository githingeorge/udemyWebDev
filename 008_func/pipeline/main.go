package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"double": double,
	"square": square,
	"sqRoot": sqRoot,
}

func double(x int) int {
	return x + x
}
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}
func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", 3)
	if err != nil {
		log.Fatal(err)
	}

}
