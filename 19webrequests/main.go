package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://deveekly.vercel.app"

func main() {
	fmt.Println("Welcome to Web Requests in Golang")

	// GET Method
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("The content from the response is: ", content)
}
