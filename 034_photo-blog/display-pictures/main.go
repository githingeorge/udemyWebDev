package main

import (
	"html/template"
	"net/http"

	"strings"

	"log"

	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/githingeorge/udemyWebDev/034_photo-blog/store-values/session"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := session.GetCookie(w, req)
	if req.Method == http.MethodPost {
		fname := handleForm(req)
		cookie = setImagesCookie(w, cookie, fname)
	}
	data := strings.Split(cookie.Value, "|")[1:]
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

func handleForm(req *http.Request) string {
	image, imageHeader, err := req.FormFile("picture")
	if err != nil {
		log.Println(err)
	}
	defer image.Close()
	fmt.Println(imageHeader.Filename)
	// create sha1 for filename
	ext := strings.Split(imageHeader.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, image)
	fname := fmt.Sprintf("%x.%s", h.Sum(nil), ext)

	// create new file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := filepath.Join(wd, "public", "pics", fname)
	newFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	image.Seek(0, 0)
	io.Copy(newFile, image)
	return fname
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
