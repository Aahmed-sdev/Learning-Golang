package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "friday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days { // i holds the index
		fmt.Println(i)
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	for _, day := range days {
		fmt.Printf("value is %v \n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		// if rougueValue == 5 {
		// 	break
		// }
		if rougueValue == 6 {
			rougueValue++
			continue
		}

		if rougueValue == 8 {
			goto lco
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}

lco: //Go to statement
	fmt.Println("Jumping to learn code online")

}
