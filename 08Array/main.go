package main

import "fmt"

func main() {
	fmt.Println("Welcome ")

	var fruitList [4]string // Compulsion to give the size

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is :", len(fruitList))
	fmt.Println("Fruit list is :", fruitList[1])

	var vegList = [3]string{"Potato", "bean", "mushroom"}
	fmt.Println("Vegy list is :", vegList)
}
