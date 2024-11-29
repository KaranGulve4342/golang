package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// For loop

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }

	for _, day := range days {
		fmt.Printf("The value is %v\n", day)
	}

	fmt.Println()

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			// break it
			rougueValue += 2
		}

		if rougueValue == 7 {
			goto newFunc
		}
		fmt.Printf("Value is %v\n", rougueValue)
		rougueValue++
	}

newFunc:
	fmt.Println("This is a new function")

}
