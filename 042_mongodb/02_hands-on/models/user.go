package models

import (
	"encoding/json"
	"os"

	"fmt"
)

const DBName = "data"

// func init() {

// 	file, err := os.Create(DBName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	users = DbUsers{}

// 	err = json.NewDecoder(file).Decode(&users)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// type DbUsers map[string]User

// func (db DbUsers) Set(user User) (User, error) {
// 	ID := uuid.NewV4()
// 	user.Id = ID.String()

// 	db[ID.String()] = user

// 	file, err := os.Open(DBName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	err = json.NewEncoder(file).Encode(db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return user, nil

// }

// var users DbUsers

// func GetSession() DbUsers {
// 	return users
// }

func StoreUsers(m map[string]User) {
	f, err := os.Create(DBName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	m := make(map[string]User)

	f, err := os.Open(DBName)
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}

	return m
}
