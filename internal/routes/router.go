package routes

import (
	"CRUD-API/internal/handlers"
	"CRUD-API/internal/repo"
	"CRUD-API/internal/services"

	"github.com/gorilla/mux"
)


func SetupRouter() *mux.Router{
	r := mux.NewRouter()

	repo.Seed()

	svc := services.NewMovieService()
	h := handlers.NewmovieHnadler()
}