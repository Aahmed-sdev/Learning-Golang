package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=frgfdr5654"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myUrl)

	//Parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//Query param
	qparam := result.Query()
	fmt.Printf("The type query params are: %T\n", qparam)

	fmt.Println(qparam["coursename"])

	for _, val := range qparam {
		fmt.Println("Param is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=azad",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
