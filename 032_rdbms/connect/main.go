package main

import (
	"database/sql"
	"log"

	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/udemy_go_web_dev?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully completed")
	check(err)
}
