package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// initialize	 data
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

func updatePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	personID := ps.ByName("id")
	found := false

	for index, item := range people {
		if item.ID == personID {
			updatedPerson := &people[index]
			updatedPerson.ID = personID
			json.NewDecoder(r.Body).Decode(&updatedPerson)
			json.NewEncoder(w).Encode(updatedPerson)
			found = true
		}
	}

	if !found {
		fmt.Fprintf(w, "No person with ID of "+personID)
	}

}

// main func
func main() {
	// initialize router
	router := httprouter.New()

	// add data
	people = append(people, person{ID: "1", Name: "Sangwoo", Age: 29, Address: &address{City: "Los Angeles", State: "CA"}})
	people = append(people, person{ID: "2", Name: "Paul", Age: 28, Address: &address{City: "Irvine", State: "CA"}})

	// routers
	router.PUT("/people/:id", updatePerson)

	// prints the message on bash
	log.Println("running api server on port 8000")

	// listens port 8000 and add router
	http.ListenAndServe(":8000", router)
}
