package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "user_gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user_gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't see all users")
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user_gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't create new user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New user created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user_gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't create new user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user_gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't create new user")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "User updated")
}
