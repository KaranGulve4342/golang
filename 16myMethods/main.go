package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	karan4342 := User{"Karan", "karanpg2306@gmail.com", true, 21, 21}

	fmt.Println(karan4342)

	fmt.Printf("Karan's Name is %v and his email is %v\n", karan4342.Name, karan4342.Email)

	fmt.Printf("Karan details is %+v\n", karan4342)

	fmt.Printf("Karan's status is %v\n", karan4342.GetStatus())

	karan4342.NewMail()

	fmt.Printf("Karan details is %+v\n", karan4342)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	OneAge int
}

func (u User) GetStatus() string {
	if u.Status {
		return "Active"
	} else {
		return "Inactive"
	}
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"

	fmt.Println("Email has been changed to", u.Email)
	fmt.Println()
}
