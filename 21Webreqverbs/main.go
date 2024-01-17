package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(content)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is :", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//Fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Lets learn go lang",
			"price": 0,
			"Platform":"lco.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "Azad")
	data.Add("LastName", "Ahmed")
	data.Add("email", "test@test.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
