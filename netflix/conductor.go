package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"Desc"`
// 	Content string `json:"Content"`
// }

type Netflix struct {
	Title string `json:"Title"`
}

type netflix []Netflix

var articles = netflix{
	Netflix{Title: "Cricket"},
}

// func listArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Inside allArticles")
// 	json.NewEncoder(w).Encode(articles)
// }

// func postArticles(w http.ResponseWriter, r *http.Request) {
// 	var newArticle Article
// 	fmt.Println("Inside getArticles")
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&newArticle)
// 	if err != nil {
// 		fmt.Println("Inside error")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println(newArticle)
// 	// body, _ := ioutil.ReadAll(r.Body)
// 	// _ = json.Unmarshal([]byte(body), &newArticle)
// 	// fmt.Println(newArticle)
// 	articles = append(articles, newArticle)
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Homepage")
}

// func getmethod(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("heloo world")
// }

func getmethod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside allArticles")
	json.NewEncoder(w).Encode(articles)
	fmt.Println("heloo world")
	// return "heloo"
}

// func putmethod(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("heloo world")
// 	var newArticle Netflix
// 	fmt.Println("Inside getArticles")
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&newArticle)
// 	if err != nil {
// 		fmt.Println("Inside error")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println(newArticle)
// 	articles = append(articles, newArticle)
// }

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/all", listArticles).Methods("GET")
	// myRouter.HandleFunc("/all", postArticles).Methods("POST")
	myRouter.HandleFunc("/conductor", getmethod).Methods("GET")
	// myRouter.HandleFunc("/conductor", putmethod).Methods("POST")
	log.Fatal(http.ListenAndServe(":9091", myRouter))
}

func main() {
	handleRequests()
}
