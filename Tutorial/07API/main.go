package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model
type Memo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	WordCount int `json:"wordcount"`
	Author *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
}

// Database
var memos []Memo

// Middleware
func (m *Memo) IsEmpty() bool {
	return /*(m.Id == "") &&*/ (m.Title == "")
}

func main () {
	fmt.Println("Starting API")
	r := mux.NewRouter()
	memos = append(memos, Memo{Id: "3", Title: "title3", WordCount: 5, Author: &Author{Fullname: "Kinantan A B", Email: "kinantanyolo@gmail.com"}})
	memos = append(memos, Memo{Id: "2", Title: "title2", WordCount: 3, Author: &Author{Fullname: "Kinantan A B", Email: "kinantanyolo@gmail.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/memos", getAllMemos).Methods("GET")
	r.HandleFunc("/memo/{id}", getOneMemo).Methods("GET")
	r.HandleFunc("/memo", createOneMemo).Methods("POST")
	r.HandleFunc("/memo/{id}", updateOneMemo).Methods("PUT")
	r.HandleFunc("/memo/{id}", deleteOneMemo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>serveHome</h1>"))
}
func getAllMemos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllMemos")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memos)
}
func getOneMemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOneMemo")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, memo := range memos {
		if memo.Id == params["id"] {
			json.NewEncoder(w).Encode(memo)
			return
		}
	}
	json.NewEncoder(w).Encode("memo Id not found")
}
func createOneMemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createOneMemo")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("no body")
		return
	}
	
	var memo Memo
	_ = json.NewDecoder(r.Body).Decode(&memo)
	if memo.IsEmpty() {
		json.NewEncoder(w).Encode("no content")
		return
	}

	rand.Seed(time.Now().UnixNano())
	memo.Id = strconv.Itoa(rand.Intn(100000))
	memos = append(memos, memo)
	json.NewEncoder(w).Encode(memo)
}
func updateOneMemo (w http.ResponseWriter, r *http.Request){
	fmt.Println("updateOneMemo")
	w.Header().Set("Content-Type", "application/json")
	/*
	if r.Body == nil {
		json.NewEncoder(w).Encode("no body")
		return
	}
	var memo Memo
	_ = json.NewDecoder(r.Body).Decode(&memo)
	if memo.IsEmpty() {
		json.NewEncoder(w).Encode("no content")
		return
	}
	*/
	params := mux.Vars(r)
	for idx, memo := range memos {
		if memo.Id == params["id"] {
			memos = append(memos[:idx], memos[idx+1:]...)
			var memo Memo
			_ = json.NewDecoder(r.Body).Decode(&memo)
			memo.Id = params["id"]
			memos = append(memos, memo)
			json.NewEncoder(w).Encode(memo)
			return
		}
	}
}

func deleteOneMemo (w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteOneMemo")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for idx, memo := range memos {
		if memo.Id == params["id"] {
			memos = append(memos[:idx], memos[idx+1:]...)
			return
		}
	}
}
