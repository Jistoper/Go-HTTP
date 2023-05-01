package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name    string `json:"Name"`
	Nim     string `json:"Nim"`
	Address string `json:"Address"`
}

var students []Student

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("Name")
		nim := r.FormValue("Nim")
		address := r.FormValue("Address")

		student := Student{
			Name:    name,
			Nim:     nim,
			Address: address,
		}

		students = append(students, student)

		json, err := json.Marshal(student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	} else if r.Method == "GET" {
		json, err := json.Marshal(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
