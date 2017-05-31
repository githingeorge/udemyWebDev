package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (hd hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "any code in this func")
}

func main() {
	var d hotDog
	http.ListenAndServe(":8000", d)

}
