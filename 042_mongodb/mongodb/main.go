package main

import (
	"net/http"

	"github.com/githingeorge/udemyWebDev/042_mongodb/mongodb/controllers"

	"log"

	"github.com/pressly/chi"
	"gopkg.in/mgo.v2"
)

func main() {
	r := chi.NewRouter()
	uc := controllers.NewUserController(getSession())
	r.Get("/", uc.Index)
	r.Get("/user/:id", uc.GetUser)
	r.Post("/user", uc.CreateUser)
	r.Delete("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}

	return s
}
