package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is: ", fruitList)

	fmt.Println("Fruit list is: ", len(fruitList))
	fmt.Println("Fruit at 3rd position is: ", fruitList[2])

	fmt.Println("Length of fruit at 3rd position is: ", len(fruitList[2]))

	// Array with values
	vegList := [3]string{"Potato", "Tomato", "Brinjal"}

	fmt.Println("Veg list is: ", vegList)

}
