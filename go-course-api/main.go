package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"slices"
	"strconv"
)

// Model for course-file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice float64 `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middleware, Helper files
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// Controllers files

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Write([]byte("<h1>Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// /courses/{id}
	id := r.PathValue("id")
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	// check Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// check Body is {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		fmt.Println("Error while decoding JSON |", err)
		return
	}

	name := course.CourseName
	fmt.Println("NAme", course.CourseName)

	// check duplicate course name
	for _, course := range courses {
		if course.CourseName == name {
			json.NewEncoder(w).Encode("Course already exist")
			return
		}
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// Generate unique id
	course.CourseId = strconv.Itoa(rand.IntN(1000000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	// course/{id}
	params := r.PathValue("id")

	for _, course := range courses {
		if course.CourseId == params {

			// Remove the existing course
			courses = slices.DeleteFunc(courses, func(course Course) bool {
				return course.CourseId == params
			})

			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				fmt.Println("Error while decoding |", err)
				return
			}
			course.CourseId = params

			// Add the updated code
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Can't found the course with given ID")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	length := len(courses)

	// course/{id}
	params := r.PathValue("id")

	for _, course := range courses {
		if course.CourseId == params {
			courses = slices.DeleteFunc(courses, func(course Course) bool {
				return course.CourseId == params
			})
		}
	}

	if length == len(courses) {
		json.NewEncoder(w).Encode("Can't found the course with given ID")
		return
	}
	json.NewEncoder(w).Encode("Course deleted successfully")
	return
}

func main() {

	http.HandleFunc("GET /", serveHome)
	http.HandleFunc("GET /courses", getAllCourses)
	http.HandleFunc("GET /course/{id}", getOneCourse)
	http.HandleFunc("POST /course", createOneCourse)
	http.HandleFunc("PUT /course/{id}", updateOneCourse)
	http.HandleFunc("DELETE /course/{id}", deleteOneCourse)

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
