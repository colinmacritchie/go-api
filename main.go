package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

type Person struct {

	ID		    string `json:"id,omitempty"`
	FrstName  string `json:"firstname, omitempty"`
	LastName  string `json:"lastname, omitempty"`
	Address   *Address `json:address, omitempty"`
}

type Address struct {
	City	string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
	
}

// Set up people lsit
var people []Person


// Main Function
func main() {
  router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", LastName: "Doe", Address: &Address{City:"City X", State: "State X"}})
  people = append(people, Person{ID: "2", Firstname: "Tedyy", LastName: "Mitchel", Address: &Address{City:"City y", State: "State y"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
