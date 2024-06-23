package main

import (
	"go-first/internal/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	log.Println("Server starting on port 9000...")
	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
