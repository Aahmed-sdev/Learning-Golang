package main

import "fmt"

func main() {
	fmt.Println("Welcome to aclass on pointers")

	// var ptr *string
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 26

	var ptr = &myNumber                              //  Reference -> &
	fmt.Println("Value of actual pointer is ", ptr)  // Actual memory address
	fmt.Println("Value of actual pointer is ", *ptr) // Actual value

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber) // operations are performed in actual value

}
