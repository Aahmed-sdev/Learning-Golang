package main

import "fmt"

func main() {
	fmt.Println("Structs in goland")

	//no inheritance in golang
	// no super or parent

	azad := User{"Azad", "aahmed@go.dev", true, 25}
	fmt.Println(azad)
	fmt.Printf("Azad details are: %+v\n", azad)
	fmt.Printf("name is %v and email is %v \n", azad.Name, azad.Email)
	azad.getStatus()
	fmt.Println("Before calling NewMail", azad.Email)
	azad.NewMail()
	fmt.Println("After calling NewEmail :", azad.Email)
	fmt.Println(azad)
}

type User struct { // User "U"
	Name   string
	Email  string
	Status bool
	Age    int
}

// Fisrt letter is in Capital to give public access

func (u User) getStatus() { // Methos is just a function inside a struct
	fmt.Println("Is user actie: ", u.Status)
}

func (usr User) NewMail() {
	usr.Email = "test@go.dev"
	fmt.Println("Email of this user is :", usr.Email)
}
