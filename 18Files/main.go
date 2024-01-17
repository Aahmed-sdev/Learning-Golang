package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a files"

	file, err := os.Create("./myGoFiles.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./myGoFiles.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("text data inside the file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
