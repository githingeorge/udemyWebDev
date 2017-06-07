package session

import (
	"net/http"
	"time"

	"github.com/githingeorge/udemyWebDev/030_sessions/expire-sessions-refactored/user"
	"github.com/satori/go.uuid"
)

type Session struct {
	username     string
	lastActivity time.Time
}

var dbSessions = map[string]Session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	dbSessionsCleaned = time.Now()
}

func get(w http.ResponseWriter, req *http.Request) Session {
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = newCookie()
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	session, ok := sessionExists(cookie.Value)

	return dbSessions[cookie.Value]

}

func sessionExists(id string) (session, bool) {
	session, ok := dbSessions[id]
	if ok {
		session.lastActivity = time.Now()
	}
	dbSessions[id] = session

	return session, ok
}

func newSession(username string) Session {
	return Session{username, time.Now()}
}

func newCookie() *http.Cookie {
	id := uuid.NewV4()
	return &http.Cookie{
		Name:  "session",
		Value: id.String(),
		// HttpOnly: true,
	}
}

func alreadyLoggedIn(req *http.Request) bool {
	// if user already exists , get user
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	return isInSession(cookie.Value)

}

func isInSession(id string) bool {
	if session, ok := dbSessions[id]; ok {
		session.lastActivity = time.Now()
		dbSessions[id] = session
		if _, ok := user.GetByUserName(id); ok {
			return ok
		}
	}
	return false
}
