package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func main() {
	fmt.Println("Golang REST Mux API")
	Articles = []Article{
		Article{Title: "Harry Potter: The Beginning", Description: "Elit excepteur amet sint labore cillum duis est.", Content: "Deserunt dolore irure esse ipsum qui nostrud ipsum velit ex non qui."},
		Article{Title: "Harry Potter: The Ned", Description: "Cupidatat magna ad adipisicing proident aliqua eu.", Content: "Eu tempor officia amet reprehenderit laborum nostrud irure deserunt commodo ad."},
	}
	handleRequests()
}
