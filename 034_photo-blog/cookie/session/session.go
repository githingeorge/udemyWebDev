package session

import (
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

const CookieName = "session"

// Session ...
type session struct {
	Username     string
	LastActivity time.Time
}

var dbSessions = map[string]session{}

const sessionLength = 60 * 10

func Get(name string) (string, session) {
	session, ok := dbSessions[name]
	if ok {
		session.LastActivity = time.Now()
	}
	return name, session
}

func GetCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
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
	return cookie
}
