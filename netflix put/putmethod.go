package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Netflix struct {
	Title string `json:"Title"`
}

type netflix []Netflix

var articles = netflix{
	Netflix{Title: "Football"},
}

func putmethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("heloo world")
	var newArticle Netflix
	fmt.Println("Inside getArticles")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newArticle)
	if err != nil {
		fmt.Println("Inside error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(newArticle)
	articles = append(articles, newArticle)
}

func handleRequests1() {

	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/all", listArticles).Methods("GET")
	// myRouter.HandleFunc("/all", postArticles).Methods("POST")
	// myRouter.HandleFunc("/conductor", getmethod).Methods("GET")
	myRouter.HandleFunc("/conductor", putmethod).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}

func main() {
	handleRequests1()
}
