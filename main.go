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

func homePage(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	welcomeMessage := `
		<div style="display:grid;height:75vh">
			<h1 style="text-align:center; margin:auto"> Welcome to SQL ANGELES </h1>
		</div>`

	fmt.Fprintf(w, welcomeMessage)
}

func getPeople(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	personID := ps.ByName("id")
	for _, item := range people {
		if item.ID == personID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf(w, "<h1>No DATA</h1>")
}

// main func
func main() {
	// initialize router
	router := httprouter.New()

	// add data
	people = append(people, person{ID: "1", Name: "Sangwoo", Age: 29, Address: &address{City: "Los Angeles", State: "CA"}})
	people = append(people, person{ID: "2", Name: "Paul", Age: 28, Address: &address{City: "Irvine", State: "CA"}})

	// routers
	router.GET("/", homePage)
	router.GET("/people", getPeople)
	router.GET("/people/:id", getPerson)

	log.Println("server is running on port localhost:8000")

	// listens port 8000 and add router
	http.ListenAndServe(":8000", router)
}
