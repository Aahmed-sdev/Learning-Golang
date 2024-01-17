package main

import "fmt"

const LoginToken string = "fsefsefsef" // Public

func main() {
	var username string = "Azad"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float64 = 255.4564654646464
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//Aliases
	// default values and aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var websit = "abvc"
	fmt.Println(websit)

	// no var  style
	numberfofUser := 30000
	fmt.Println(numberfofUser)

	fmt.Println(LoginToken)

}
