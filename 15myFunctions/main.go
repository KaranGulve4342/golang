package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in Go")
	greeterTow()

	result := adder(3, 5)

	fmt.Println("The sum of 3 and 5 is", result)

	proresult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("The reuslt of ProAdder is", proresult)
}

func greeter() {
	fmt.Println("Hello, World!")
}
func greeterTow() {
	fmt.Println("Its is a greeter two function")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}
