package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jahmin/golang-restapi/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	log.Println("Server started on :28082")
	log.Fatal(http.ListenAndServe(":28082", r))
}

