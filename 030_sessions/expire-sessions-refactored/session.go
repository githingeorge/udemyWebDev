package main

import (
	"net/http"

	"fmt"

	"time"

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
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	// if user already exists , get user
	var user User
	if session, ok := dbSessions[cookie.Value]; ok {
		session.lastActivity = time.Now()
		dbSessions[cookie.Value] = session
		user = dbUsers[session.username]
	}

	return user

}

func alreadyLoggedIn(req *http.Request) bool {
	// if user already exists , get user
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	if session, ok := dbSessions[cookie.Value]; ok {
		session.lastActivity = time.Now()
		dbSessions[cookie.Value] = session
		if _, ok := dbUsers[session.username]; ok {
			return ok
		}
	}
	return false
}

func cleanSessions() {
	fmt.Println("Before Clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second*30) || v.username == "" {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After clean")
	showSessions()
}

func showSessions() {
	fmt.Println("-------")
	for k, v := range dbSessions {
		fmt.Println(k, v.username)
	}
	fmt.Println("---")
}
