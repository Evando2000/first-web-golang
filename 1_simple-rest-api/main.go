package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test title",
			Desc:    "Test Desc",
			Content: "Hello world"},
		Article{
			Title:   "Test title",
			Desc:    "Test Desc",
			Content: "Hello world"}}
	fmt.Println("Endpoint all articles")
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
