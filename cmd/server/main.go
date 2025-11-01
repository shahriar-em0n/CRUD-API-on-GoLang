package main

import (
	"CRUD-API/internal/repo"
	"CRUD-API/internal/routes"
	"log"
	"net/http"
)

func main() {
	repo.SeedData()

	router := routes.SetupRouter()

	log.Println(" Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
