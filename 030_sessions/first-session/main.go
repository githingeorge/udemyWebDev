package main

import (
	"html/template"
	"net/http"

	"fmt"

	"log"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}

	// if user already exists , get user
	var user User
	if username, ok := dbSessions[cookie.Value]; ok {
		user = dbUsers[username]
	}

	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		user = User{
			username,
			req.FormValue("firstname"),
			req.FormValue("lastname"),
		}
		dbSessions[cookie.Value] = username
		dbUsers[username] = user
	}
	fmt.Println("index: ", dbSessions, dbUsers)

	err = tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("bar: ")
	fmt.Println("bar: ", dbSessions, dbUsers)

	c, err := req.Cookie("session")
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println(dbSessions)
	username, ok := dbSessions[c.Value]
	if !ok {
		log.Fatal(err)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	user := dbUsers[username]
	tpl.ExecuteTemplate(w, "bar.gohtml", user)
}

// 4f855aff-e05d-45a5-9830-25cf896bf264
