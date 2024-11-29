package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// No Inheritance in Golang; No super or parent

	karan := User{"Karan", "karanpg2306@gmail.com", true, 20}

	fmt.Println(karan)

	// print the fields of the struct
	fmt.Printf("Karan details are: %+v\n", karan)                        // %+v will print the field names as well
	fmt.Printf("Name is %v and Email is %v.\n", karan.Name, karan.Email) // accessing fields using the dot operator

}

// Structs are similar to classes in OOP languages
// Structs are value types
// Structs are used to create custom data types
// Structs can have fields and methods
// Structs can have anonymous fields
// Structs can be embedded in other structs
// Structs can implement interfaces
// Structs are used to create reusable code
// Structs are used to create custom data types

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
