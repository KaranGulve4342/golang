package main

import "fmt"

func main() {
	// Code here
	fmt.Println("Welcome to Maps in Go!")

	fmt.Println()
	// Creating a map
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	languages["TS"] = "TypeScript"
	languages["CPP"] = "C++"

	fmt.Println("List of all Languages:", languages)

	fmt.Println("JS is:", languages["JS"])

	fmt.Println()

	delete(languages, "RB")

	fmt.Println("List of all Languages:", languages)

	// Loops in Maps

	fmt.Println()

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	fmt.Println()
}
