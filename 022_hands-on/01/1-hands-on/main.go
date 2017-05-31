package main

import "net/http"
import "fmt"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}
func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "dog")

}
func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "githin")

}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8000", nil)
}
