package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("Present time is: ", presentTime)
	fmt.Println("Present time in format is: ", presentTime.Format(("01-02-2006")))                 // MM-DD-YYYY
	fmt.Println("Present time in format is: ", presentTime.Format(("02-01-2006 15:04:05")))        // DD-MM-YYYY HH:MM:SS
	fmt.Println("Present time in format is: ", presentTime.Format(("02-01-2006 15:04:05 Monday"))) // DD-MM-YYYY HH:MM:SS Day

	createdDate := time.Date(2004, time.June, 23, 4, 4, 0, 0, time.Now().Location())

	fmt.Println("Created date is: ", createdDate)

	fmt.Println("created date in format is: ", createdDate.Format(("02-01-2006 Monday"))) // DD-MM-YYYY Day
}

// GOOS
// use of GOOS environment variable to get the operating system name
// cmd : go env GOOS

// $env:GOOS="windows"; go build main.go // windows
// $env:GOOS="linux"; go build main.go // linux
// $env:GOOS="darwin"; go build main.go // mac

// example :
// GOOS=linux go build main.go // linux

// GOOS=windows go build main.go // windows
// GOOS=darwin go build main.go // mac

// why to build a program for multiple operating systems
//  To provide a better user experience

// GOARCH
// use of GOARCH environment variable to get the architecture name

// cmd : go env GOARCH
// ex : amd64
