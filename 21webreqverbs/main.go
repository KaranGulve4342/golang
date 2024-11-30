package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb in Golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	fmt.Println("Performing GET Request")

	const myURL = "http://localhost:3000/get"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.Status)
	fmt.Println("Content length is:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// fmt.Println("Content is:", content) // 1st method
	// fmt.Println("Content is:", string(content))

	var responseString strings.Builder // 2nd method

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is written:", byteCount)

	fmt.Println("Content is:", responseString.String())

	// Why 2nd method is better? =>
	// Because it is more efficient
	// as it doesn't create a new string every time we append to it
	// unlike the 1st method

}
