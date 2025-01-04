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

type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"auther"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

//middleware, helper - usually go in seperate file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API for courses")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId: "lajsdlkjf", CourseName: "java", Price: 333, Author: &Author{Fullname: "abc xyz", Website: "google.com"}})
	courses = append(courses, Course{
		CourseId: "adsfasdf", CourseName: "Python", Price: 785, Author: &Author{Fullname: "abc xyz", Website: "google.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - usually go in seperate file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	//grab all params from request
	params := mux.Vars(r)

	//find course with id idParam
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saving a course")
	w.Header().Set("Content-type", "application/json")

	// if body of request is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid json")
	}
	// if body is {}

	//checking validity if json

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send valid data/json")
	}

	//generate a unique id, convert it to string
	//append new course to courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a course")
	w.Header().Set("Content-type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	//loop through courses, remove that course, add updated course
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//todo : send a response when id is not found
	json.NewEncoder(w).Encode("course with given id not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting a course")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses to find course
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("course with given id not found")
}
