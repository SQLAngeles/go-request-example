package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// initialize data
type person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *address `json:"address,omitempty"`
}
type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []person

func createPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newID := ps.ByName("id")
	var newPerson person
	json.NewDecoder(r.Body).Decode(&newPerson)
	newPerson.ID = string(newID)
	people = append(people, newPerson)
	json.NewEncoder(w).Encode(people)
}

// main func
func main() {
	// initialize router
	router := httprouter.New()

	// add data
	people = append(people, person{ID: "1", Name: "Sangwoo", Age: 29, Address: &address{City: "Los Angeles", State: "CA"}})
	people = append(people, person{ID: "2", Name: "Paul", Age: 28, Address: &address{City: "Irvine", State: "CA"}})

	// routers
	router.POST("/people/:id", createPerson)

	log.Println("server is running on port localhost:8000")

	// listens port 8000 and add router
	http.ListenAndServe(":8000", router)
}
