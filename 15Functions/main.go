package main

import "fmt"

func main() {
	fmt.Println("Welcome to go function")
	greeter()
	// func greeterTwo(){   //this is not allowed
	// 	fmt.Println("Function two")
	// }
	greeterTwo()

	result := adder(3, 4)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is:", proResult)
	fmt.Println(":Pro Message is:", myMessage)
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func greeterTwo() {
	fmt.Println("Namaste from golang TWO")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}
