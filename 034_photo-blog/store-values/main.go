package main

import (
	"html/template"
	"net/http"

	"strings"

	"github.com/githingeorge/udemyWebDev/034_photo-blog/store-values/session"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := session.GetCookie(w, req)
	cookie = setImagesCookie(w, cookie, "disney.jpg", "sunset.jpg", "beach.jpg")
	data := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

func setImagesCookie(w http.ResponseWriter, cookie *http.Cookie, images ...string) *http.Cookie {
	c := cookie.Value
	for _, image := range images {
		if !strings.Contains(c, image) {
			c += "|" + image
		}
	}
	cookie.Value = c

	http.SetCookie(w, cookie)
	return cookie
}
