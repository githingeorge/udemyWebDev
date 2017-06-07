package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) User {
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

	return user

}

func alreadyLoggedIn(req *http.Request) bool {
	// if user already exists , get user
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	if username, ok := dbSessions[cookie.Value]; ok {
		if _, ok := dbUsers[username]; ok {
			return ok
		}
	}
	return false
}
