package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// models -- file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middlewares -- file
func (c *Course) IsEmpty() bool { // declaring a function like this makes it exclusive to the course struct
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers -- file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>You talkin' to me?</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// extract params(course id) from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through fake db to find the course matching the course id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate!!")
			return
		}
	}

	// generate unique id string
	// append course in courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id , remove, add with new id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted the course successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func main() {

	fmt.Println("API Tutorial")
	r := mux.NewRouter()

	// seeding with fake data
	courses = append(courses, Course{CourseId: "1", CourseName: "Life",
		CoursePrice: "Infinite", Author: &Author{FullName: "Charles Bukowski", Website: "Don't have one"},
	})
	courses = append(courses, Course{CourseId: "2", CourseName: "Dictatorship",
		CoursePrice: "Zero", Author: &Author{FullName: "George Orwell", Website: "They blocked it!"},
	})
	courses = append(courses, Course{CourseId: "3", CourseName: "Piracy",
		CoursePrice: "A million bucks and a few shanties", Author: &Author{FullName: "Captain Jack Sparrow", Website: "On the way darling;)"},
	})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	// listen and serve on a port
	log.Fatal(http.ListenAndServe(":4000", r))

}
