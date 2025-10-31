package handlers

import (
	"CRUD-API/internal/models"
	"CRUD-API/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


type MovieHandler struct{
	svc *services.MovieService
}

func NewmovieHnadler(svc *services.MovieService) *MovieHandler{
	return &MovieHandler{svc: svc}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request){
	response.JSON(w, http.StatusOK, h.svc.List())
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	m, err := h.svc.Get(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, m)
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	created, err := h.svc.Create(m)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, created)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var m models.Movie
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	updated, err := h.svc.Update(id, m)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, updated)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	remaining, _ := h.svc.Delete(id)
	response.JSON(w, http.StatusOK, remaining)
}
