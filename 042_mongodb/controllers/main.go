package main

import (
	"net/http"

	"github.com/githingeorge/udemyWebDev/042_mongodb/controllers/controllers"

	"log"

	"github.com/pressly/chi"
)

func main() {
	r := chi.NewRouter()
	uc := controllers.NewUserController()
	r.Get("/", uc.Index)
	r.Get("/user/:id", uc.GetUser)
	r.Post("/user", uc.CreateUser)
	r.Delete("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
