package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", GetUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	InitialMigration()
	handleRequests()
}
