package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File-Handling in Golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		fmt.Println("Error creating file")
		panic(err) // stop the program if file creation failed
	} else {
		fmt.Println("File created successfully")
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Error writing to file")
		panic(err) // stop the program if file writing failed
	} else {
		fmt.Println("File written to successfully with length", length)
	}

	defer file.Close() // recommended to close the file after the work is done

	readFile("./myFile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data read from file is:", (databyte))
	fmt.Println("Data read from file is:", string(databyte))
}
