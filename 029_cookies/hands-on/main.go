package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func set(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visits")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Visits: ", count)
}

// func read(w http.ResponseWriter, req *http.Request) {
// 	c, err := req.Cookie("my-cookie")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	fmt.Fprintln(w, "Your Cookie:", c)
// }
