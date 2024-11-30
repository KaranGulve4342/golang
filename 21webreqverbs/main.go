package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb in Golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	fmt.Println("Performing POST Request")

	const myURL = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"name" : "karan",
			"age" : "20",
			"email" : "karanpg2306@gmail.com",
			"phone" : 1234567890,
			"address" : "India"
		}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content is:", string(content))
	fmt.Println("Content is:", content)
}

func PerformPostFormRequest() {
	fmt.Println("Performing POST form Request")

	// formdata
	const myURL = "http://localhost:3000/postform"

	data := url.Values{}

	data.Add("name", "karan")
	data.Add("age", "20")
	data.Add("email", "karanpg2306@gmail.com")
	data.Add("phone", "1234567890")
	data.Add("address", "India")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content is:", string(content))
	fmt.Println("Content is:", content)
}
