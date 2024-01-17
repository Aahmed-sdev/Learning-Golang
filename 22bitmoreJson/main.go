package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "learncodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeOnline.in", "azad123", nil},
	}

	//package this data as JSON data
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN",
		"Price": 199,
		"Platform": "learncodeOnline.in",
		"Password": "bcd123",
		"Tags": ["full-stack","js"]
	}
	`)

	var lcCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcCourse)
		fmt.Printf("%#v\n", lcCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlinedata map[string]interface{} // value is not gurantee type
	json.Unmarshal(jsonDataFromWeb, &myOnlinedata)
	fmt.Printf("%#v\n", myOnlinedata)

	for k, v := range myOnlinedata {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
