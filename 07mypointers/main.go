package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	var a int = 42
	var b *int = &a

	fmt.Println(a, b)
	fmt.Println(&a, *b)

	// var ptr *int // default value is <nil>
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of myNumber is", myNumber)
	fmt.Println("Address of myNumber is", &myNumber)
	fmt.Println("Value of myNumber is", *(&myNumber))

	fmt.Println("ptr :", ptr)
	fmt.Println("&ptr :", &ptr)
	fmt.Println("*ptr :", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value of *ptr is", *ptr)

}
