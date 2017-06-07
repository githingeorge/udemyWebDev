package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"

	"github.com/githingeorge/udemyWebDev/042_mongodb/mongodb/models"
	"github.com/pressly/chi"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	fmt.Println(id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oId := bson.ObjectIdHex(id)
	fmt.Println(oId)

	user := models.User{}

	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oId).One(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(user)

	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 201
	fmt.Fprintf(w, "%s\n", j)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request) {
	user := models.User{}

	json.NewDecoder(req.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	if err := uc.session.DB("go-web-dev-db").C("users").Insert(user); err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oId := bson.ObjectIdHex(id)

	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oId); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "deleted user", oId, "\n")
}

// 5934287d4e70bb3294f7a5a8
