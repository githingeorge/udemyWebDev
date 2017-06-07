package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/githingeorge/udemyWebDev/042_mongodb/hands-on-1/models"
	"github.com/pressly/chi"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController() *UserController {
	return &UserController{
		map[string]models.User{},
	}
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

	user, ok := uc.session[id]
	if !ok {
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

	ID := uuid.NewV4()
	user.Id = ID.String()

	uc.session[ID.String()] = user

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

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "deleted user", id, "\n")
}

// 5934287d4e70bb3294f7a5a8
