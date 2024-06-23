package userController

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Idade uint8  `json:"idade" validate:"required,gt=0"`
}

var users = []User{
	{Name: "Ana", Idade: 20},
	{Name: "Maria", Idade: 21},
	{Name: "Carlos", Idade: 22},
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
	validate := validator.New()

	if err := validate.Struct(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Name == newUser.Name {
			http.Error(w, "Name must be unique", http.StatusBadRequest)
			return
		}
	}

	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
