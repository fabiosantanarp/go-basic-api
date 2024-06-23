package routes

import (
	userController "go-first/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func jsonMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(jsonMiddlware)
	r.HandleFunc("/users", userController.List).Methods("GET")
	r.HandleFunc("/users", userController.Create).Methods("POST")
	return r
}
