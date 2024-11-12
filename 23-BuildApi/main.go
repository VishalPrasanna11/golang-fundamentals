package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake Db

var courses []Course

//middleware, helper -file

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
}

func main() {
	// Code;
	fmt.Println("Hello, from Vishal Prasanna's GoLang Playground!")
	fmt.Println("We can use go help to get help on go commands.")

	courses = append(courses, Course{CourseId: "1", CourseName: "GoLang", CoursePrice: 100, Author: &Author{FullName: "Vishal Prasanna", Website: "https://vishalprasanna.com"}})

	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/create", createCourseHandler).Methods("POST")
	r.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteCourse).Methods("DELETE")

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", r)
}

//Controller

func createCourse(course Course) Course {
	fmt.Println("Create Course")
	courses = append(courses, course)
	return course
}

//Create Course Handler

func createCourseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course Handler")
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	course = createCourse(course)
	json.NewEncoder(w).Encode(course)
}

//Serve Home Route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course By Id")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the url
	params := mux.Vars(r)

	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the url
	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id from the url
	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}
