package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	//defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()

}

// world, One, Two   -> Two, one, world ( LIFO )

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// world, One, Two
// 0, 1, 2, 3, 4
// hello

// -> Hello 43210, two, One, world
