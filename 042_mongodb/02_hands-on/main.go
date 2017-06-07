package main

import (
	"fmt"
	"net/http"

	"github.com/githingeorge/udemyWebDev/042_mongodb/02_hands-on/controllers"

	"github.com/pressly/chi"
)

func main() {
	fmt.Println("debug")

	r := chi.NewRouter()
	uc := controllers.NewUserController()
	r.Get("/", uc.Index)
	r.Get("/user/:id", uc.GetUser)
	r.Post("/user", uc.CreateUser)
	r.Delete("/user/:id", uc.DeleteUser)
	fmt.Println("debug")

	http.ListenAndServe(":8080", r)
}

// func getSession() *mgo.Session {
// 	s, err := mgo.Dial("mongodb://localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return s
// }
