package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple", "tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	//fruitList = append(fruitList[1:3]) // startIndex , endIndex ( non inclusive)
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	//highScore[4] = 777 //error index out bound

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) //false
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) //true

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
