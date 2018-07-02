package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Router
	router := mux.NewRouter()
	// router.HandleFunc("/people", GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
