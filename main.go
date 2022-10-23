package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := io.ReadAll(r.Body)
	var updateArticle Article
	json.Unmarshal(reqBody, &updateArticle)

	for index, article := range Articles {
		if article.Id == key {
			article.Title = updateArticle.Title
			article.Description = updateArticle.Description
			article.Content = updateArticle.Content
			Articles[index] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Golang REST Mux API")
	Articles = []Article{
		{Id: "1", Title: "Harry Potter: The Beginning", Description: "Elit excepteur amet sint labore cillum duis est.", Content: "Deserunt dolore irure esse ipsum qui nostrud ipsum velit ex non qui."},
		{Id: "2", Title: "Harry Potter: The Ned", Description: "Cupidatat magna ad adipisicing proident aliqua eu.", Content: "Eu tempor officia amet reprehenderit laborum nostrud irure deserunt commodo ad."},
	}
	handleRequests()
}
