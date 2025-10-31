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
	h := handlers.NewmovieHnadler(svc)

	// regiester routes

	r.HandleFunc("/movie", h.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", h.GetMovie).Methods("GET")
	r.HandleFunc("/movie", h.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", h.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", h.DeleteMovie).Methods("DELETE")

	return  r
}

