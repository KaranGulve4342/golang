package main

import "fmt"

// not allowed outside of method walarus operator
// jwtToken := 2000000 // error

const LoginToken string = "karan4342" // public variable = Putting first letter as Capital

func main() {
	var username string = "karan4342"
	fmt.Println(username)
	fmt.Printf("The datatype of the variable is: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("The datatype of the variable is: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The datatype of the variable is: %T \n", smallVal)

	var integerVal int64 = 9223372036854775807 // maximum value for int64
	fmt.Println(integerVal)
	fmt.Printf("The datatype of the variable is: %T \n", integerVal)

	var smallFloat float32 = 167.23456897554984
	fmt.Println(smallFloat)
	fmt.Printf("The datatype of the variable is: %T \n", smallFloat)

	var bigFloat float64 = 167.23456897554984
	fmt.Println(bigFloat)
	fmt.Printf("The datatype of the variable is: %T \n", bigFloat)

	// Alias for a datatype

	fmt.Println("\n\nAlias for a datatype")
	fmt.Println()

	// Default value of a variable
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("The datatype of the variable is: %T \n", anotherVar)

	// inplicit type
	var website = "karan4342.github.io"
	fmt.Println(website)
	fmt.Printf("The datatype of the variable is: %T \n", website)

	// no var style

	// walarus operator = allowed at local level
	numberOfUsers := 30000000.01334
	fmt.Println(numberOfUsers)
	fmt.Printf("The datatype of the variable is: %T \n", numberOfUsers)

	// Public
	fmt.Println("Login Token is:", LoginToken)
	fmt.Printf("The datatype of the variable is: %T \n", LoginToken)
}
