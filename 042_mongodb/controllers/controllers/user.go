package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/githingeorge/udemyWebDev/042_mongodb/controllers/models"
	"github.com/pressly/chi"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
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
	user := models.User{
		Name:   "James Bond",
		Gender: "MALE",
		Age:    42,
		Id:     id,
	}
	j, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", j)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request) {
	user := models.User{}

	json.NewDecoder(req.Body).Decode(&user)

	user.Id = "007"

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
