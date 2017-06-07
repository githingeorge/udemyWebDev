package main

import (
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/githingeorge/udemyWebDev/042_mongodb/json/models"
	"github.com/pressly/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", index)
	r.Get("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}
func index(w http.ResponseWriter, req *http.Request) {
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

func getUser(w http.ResponseWriter, req *http.Request) {
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
