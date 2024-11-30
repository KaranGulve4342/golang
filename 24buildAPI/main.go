package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// Model for Author

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper functions - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""

	return c.CourseId == "" || c.CourseName == "" || c.CoursePrice == 0 || c.Author == nil || c.Author.Fullname == "" || c.Author.Website == ""
}

func main() {

}

// controllers - file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in Golang"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	id := params["id"]

	// loop through courses to find the course with the id
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")

	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data provided - Please send some data")
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Node data provided - Please provide some data")
	}

	// generate unique id, convert that id into string
	// append this new course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	id := params["id"]

	// loop through courses to find the course with the id
	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
		}
	}

	// TODO: send a response when id is not found
	json.NewEncoder(w).Encode("No Course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	id := params["id"]

	// loop through courses to find the course with the id

	for index, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
}
