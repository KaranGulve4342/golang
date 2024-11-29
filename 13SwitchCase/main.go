package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang")

	// Switch Case

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of Dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and You can open")
	case 2:
		fmt.Println("Dice Value is 2 and You can move 2 steps")
	case 3:
		fmt.Println("Dice Value is 3 and You can move 3 steps")
		fallthrough // use of fallthrough keyword = it will execute the next case also
	case 4:
		fmt.Println("Dice Value is 4 and You can move 4 steps")
	case 5:
		fmt.Println("Dice Value is 5 and You can move 5 steps")
		fallthrough
	case 6:
		fmt.Println("Dice Value is 6 and You can move 6 steps and roll the dice again")
	default:
		fmt.Println("Invalid Dice Value")
	}

}
