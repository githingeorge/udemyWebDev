package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some-value",
	})
	fmt.Fprint(w, "COOKIE WRITTEN - check your browser")
	fmt.Fprint(w, "in chrome go to: dev tools / application /cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Your Cookie:", c1)

	c2, err := req.Cookie("general")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Your Cookie:", c2)

	c3, err := req.Cookie("specific")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Your Cookie:", c3)
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "genarla cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other specific cookie",
	})

	fmt.Fprintln(w, "Cookies wriiten")
}
