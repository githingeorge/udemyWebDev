package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (hd hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("any code in this func")
}

func main() {

}
