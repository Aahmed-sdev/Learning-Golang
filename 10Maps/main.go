package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in go")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages : ", languages)
	fmt.Println(languages["JS"])
	fmt.Println(languages["NA"])

	// delete by key
	delete(languages, "RB")
	fmt.Println(languages)

	// loop through map
	for _, value := range languages {
		fmt.Printf(" value is %v \n", value)
	}

}
