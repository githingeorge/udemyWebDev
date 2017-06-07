package main

import (
	"html/template"
	"net/http"

	"github.com/githingeorge/udemyWebDev/030_sessions/expire-sessions-refactored/session"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/bar", bar)
	// http.HandleFunc("/signup", signup)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	session := session.get()
	user := getUser(session.username)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

// func bar(w http.ResponseWriter, req *http.Request) {
// 	user := getUser(w, req)
// 	if !alreadyLoggedIn(req) {
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}
// 	if user.Role != "007" {
// 		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
// 		return
// 	}
// 	tpl.ExecuteTemplate(w, "bar.gohtml", user)
// }

// func signup(w http.ResponseWriter, req *http.Request) {
// 	if alreadyLoggedIn(req) {
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 	}

// 	// process form submission
// 	if req.Method == http.MethodPost {
// 		username := req.FormValue("username")
// 		password := req.FormValue("password")
// 		firstname := req.FormValue("firstname")
// 		lastname := req.FormValue("lastname")
// 		role := req.FormValue("role")

// 		// Username taken?
// 		if _, ok := dbUsers[username]; ok {
// 			http.Error(w, "Username already taken", http.StatusForbidden)
// 		}

// 		// create session
// 		sID := uuid.NewV4()
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(w, cookie)
// 		dbSessions[cookie.Value] = Session{username, time.Now()}

// 		// store users in dbUsers
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
// 		if err != nil {
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			return
// 		}
// 		newUser := User{
// 			UserName: username,
// 			Password: hashedPassword,
// 			First:    firstname,
// 			Last:     lastname,
// 			Role:     role,
// 		}
// 		dbUsers[username] = newUser

// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}

// 	tpl.ExecuteTemplate(w, "signup.gohtml", nil)

// }

// func login(w http.ResponseWriter, req *http.Request) {
// 	if alreadyLoggedIn(req) {
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}
// 	if req.Method == http.MethodPost {

// 		var user User
// 		// process form submission
// 		username := req.FormValue("username")
// 		password := req.FormValue("password")

// 		// is there a username
// 		user, ok := dbUsers[username]
// 		if !ok {
// 			http.Error(w, "Username and/or Passwrod do not match", http.StatusForbidden)
// 			return
// 		}

// 		// does the passwords match
// 		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 		if err != nil {
// 			http.Error(w, "Username and/or Passwrod do not match", http.StatusForbidden)
// 			return
// 		}
// 		// create session
// 		sID := uuid.NewV4()
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(w, cookie)
// 		dbSessions[cookie.Value] = Session{username, time.Now()}
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}
// 	tpl.ExecuteTemplate(w, "login.gohtml", nil)
// }

// func logout(w http.ResponseWriter, req *http.Request) {
// 	if !alreadyLoggedIn(req) {
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}

// 	cookie, _ := req.Cookie("session")
// 	// remove from session
// 	delete(dbSessions, cookie.Value)

// 	// remove the cookie
// 	c := &http.Cookie{
// 		Name:   "session",
// 		Value:  "",
// 		MaxAge: -1,
// 	}
// 	http.SetCookie(w, c)

// 	// clean up dbSession
// 	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
// 		go cleanSessions()
// 	}

// 	http.Redirect(w, req, "/login", http.StatusSeeOther)

// }
