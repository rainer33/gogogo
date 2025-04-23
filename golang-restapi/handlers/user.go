package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Jahmin"},
	{ID: 2, Name: "카카"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = len(users) + 1
	users = append(users, u)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

