package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	"net/url"
)

const googleUrl string = "https://google.com?";
const dummyUrl string = "https://localhost:8000/kiri/kanan?atas=utara&bawah=selatan&judul=mata%20angin"

func main() {
	/*
	response, _ := http.Get(googleUrl);
	fmt.Printf("Type response: %T\n", response);
	defer response.Body.Close();
	
	data, _ := ioutil.ReadAll(response.Body);
	content := string(data)
	fmt.Println(content)
	*/
	
	//parsing
	parsed, _ := url.Parse(dummyUrl);
	fmt.Println(parsed.Scheme)
	fmt.Println(parsed.Host)
	fmt.Println(parsed.Path)
	fmt.Println(parsed.Port())
	fmt.Println(parsed.RawQuery)

	params := parsed.Query()
	fmt.Printf("Params type: %T\n", params)
	for key, val := range params {
		fmt.Println(key, ":", val);
	}
	urlComponents := &url.URL{
		Scheme: "https",
		Host: "localhost",
		Path: "/test",
		RawPath: "user=John",
	}
	fmt.Println(urlComponents.String())
}