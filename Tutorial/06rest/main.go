package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"nama"`
	Age int `json:"-"`
	Educations []string `json:"edu,omitempty"`
}

func EncodeJson(){
	users := []User{
		{"Kinantan A B", 22, []string{"SMA TN", "ITB"}},
		{"Pikatan A B", 20, []string{"Semesta", "UI"}},
	}
	usersJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", usersJson)
}

func DecodeJson(){
	jsonData := []byte(`{
		"nama": "Bintang M S",
		"Age": 24,
		"edu": ["Jogja", "UNSOED"]
	}`)
	var user User
	isValid := json.Valid(jsonData)
	if isValid {
		json.Unmarshal(jsonData, &user)
		fmt.Printf("%#v\n", user)
	}else{
		fmt.Println("Json not valid")
	}

	var mapJson map[string]interface{}
	json.Unmarshal(jsonData, &mapJson)
	fmt.Printf("%#v\n", mapJson)
	for k, v := range mapJson {
		fmt.Printf("%v: %v (%T)\n", k, v, v)
	}
}

func main() {
	//PerformGetRequest()
	//PerformPostJSONRequest()
	//PerformPostFormRequest()
	//EncodeJson()
	//DecodeJson()
	fmt.Println("main")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter(){
	fmt.Println("greeter")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>My learning on Golang</h1>"))
}


func PerformGetRequest() {
	const serverurl = "http://localhost:8000/get"
	response, err := http.Get(serverurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const serverurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"firstName": "Kinantan",
			"middleName": "Arya"
			"lastName": "Bagaspati"
		}
	`)

	response, err := http.Post(serverurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const serverurl = "http://localhost:8000/form"

	formData := url.Values{}
	formData.Add("firstName", "Kinantan")
	formData.Add("middleName", "Arya")
	formData.Add("lastName", "Bagaspati")

	response, err := http.PostForm(serverurl, formData)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}