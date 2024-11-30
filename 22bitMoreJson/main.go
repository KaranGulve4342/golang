package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      // if price is 0, it will not be printed
	Platform string   `json:"website"`        // if website is nil, it will not be printed
	Password string   `json:"-"`              // password will not be printed
	Tags     []string `json:"tags,omitempty"` // if tags is nil, it will not be printed
}

func main() {
	fmt.Println("Welcome to JSON in Golang!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	// Create a course

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abc@123", nil},
		{"data Science Bootcamp", 399, "LearnCodeOnline.in", "abc@123", []string{"data", "ai"}},
		{"Cyber Security Bootcamp", 499, "LearnCodeOnline.in", "abc@123", []string{"security", "data"}},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "prefix", "\t")

	if err != nil {
		panic(err)
	}

	// Print the JSON data
	// fmt.Println("JSON data: ", finalJson)
	fmt.Printf("%s\n", finalJson)
	// fmt.Println(string(finalJson))
}

func DecodeJson() {
	// json data
	jsonData := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev", "js"]
		}
	`)

	var myCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was not Valid")
	}

	// some cases where you just want to add data to eky value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	fmt.Println()

	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v and Value is: %v and Type is: %T\n", k, v, v)
	}

}
