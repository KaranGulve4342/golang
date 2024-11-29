package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	var fruitList = []string{"Apple", "Tomato", "Banana"}

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Length of fruitList is: ", len(fruitList))

	fruitList = append(fruitList, "Mango", "Peach")

	fmt.Println("Fruit List after adding Mango and Peach: ", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Length of fruitList is: ", len(fruitList))

	// fruitList = append(fruitList[1:3])

	// fruitList = append(fruitList[1:])

	// fruitList = append(fruitList[:2])

	fruitList = append(fruitList[:2], fruitList[3:]...)

	fmt.Println("Fruit List after slicing: ", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Length of fruitList is: ", len(fruitList))

	fmt.Println()

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Type of highScores is %T\n", highScores)
	fmt.Println("Length of highScores is: ", len(highScores))

	highScores = append(highScores, 555, 666, 777, 888)

	fmt.Println("High Scores after adding more: ", highScores)
	fmt.Printf("Type of highScores is %T\n", highScores)
	fmt.Println("Length of highScores is: ", len(highScores))

	fmt.Println()

	// sorting the array
	sort.Ints(highScores)

	fmt.Println("Sorted High Scores: ", highScores)

	fmt.Println("Is the array sorted: ", sort.IntsAreSorted(highScores))

	fmt.Println()

	// How to remove a value from slice based on Index

	var courses = []string{"React", "Python", "JavaScript", "Machine Learning", "Django"}

	fmt.Println("Courses: ", courses)

	var index int = 2
	// Removing the value at index 2 from the slice
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Courses after removing: ", courses)
}
