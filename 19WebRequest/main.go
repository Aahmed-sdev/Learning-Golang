package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web Request")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // callers responsibility to close the connection

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
