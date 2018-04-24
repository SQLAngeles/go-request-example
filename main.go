package main

import (
	"encoding/json"
	"fmt"
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

func deletePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	personID := ps.ByName("id")
	found := false
	for index, item := range people {
		if item.ID == personID {
			people = append(people[:index], people[index+1:]...)
			found = true
		}
	}

	if !found {
		fmt.Fprintf(w, "No person with ID of "+personID)
		return
	}
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
	router.DELETE("/people/:id", deletePerson)

	log.Println("server is running on port localhost:8000")

	// listens port 8000 and add router
	http.ListenAndServe(":8000", router)
}
