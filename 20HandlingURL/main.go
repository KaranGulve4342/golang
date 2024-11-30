package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://deveekly.vercel.app/team"

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	fmt.Println(myURL)

	// Parsing the URL

	result, _ := url.Parse(myURL)

	fmt.Println("Result after parsing the URL is :", result)
	fmt.Println("Scheme after parsing the URL is :", result.Scheme)
	fmt.Println("Host after parsing the URL is :", result.Host)
	fmt.Println("Port after parsing the URL is :", result.Port())
	fmt.Println("RawQuery after parsing the URL is :", result.RawQuery)

	qparams := result.Query()

	fmt.Println("Query parameters is :", qparams)
	fmt.Printf("The Type of Query params are: %T\n", qparams)

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=karan",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
