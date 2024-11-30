// Defer statements :
// 1. A defer statement defers the execution of a function until the surrounding function returns.
// 2. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// 3. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed

// What is the use of defer in one line?
// => Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
// => real world Example : Closing a file after opening it.

package main

import "fmt"

func main() {
	defer fmt.Println("Hello World")
	defer fmt.Println("Hello World 2")
	fmt.Println("start1")
	defer fmt.Println("Hello World 3")
	fmt.Println("start2")

	fmt.Println()

	myDefer()

}

func myDefer() {
	for i := 0; i < 7; i++ {
		defer fmt.Println(i)
	}
}
