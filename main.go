package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempry"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Do", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Mike", Lastname: "Co", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Ben", Lastname: "Go", Address: &Address{City: "City X", State: "State Z"}})
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetPeople : GET /people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {}
