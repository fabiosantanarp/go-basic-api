package userController

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Ana"},
	{ID: 2, Name: "Maria"},
	{ID: 3, Name: "Carlos"},
}

func List(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "error on serialize data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "error on decoding request", http.StatusBadRequest)
		return
	}
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
