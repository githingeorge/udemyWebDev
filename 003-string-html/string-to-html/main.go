package main

import (
	"fmt"
)

func main() {
	name := "Todd Mcleaod"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
