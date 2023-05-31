package main

import (
	"encoding/json"
	"net/http"
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
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

// controllers -- file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>You talkin' to me?</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func main() {

}
