package services

import (
	"CRUD-API/internal/models"
	"CRUD-API/internal/repo"
	"errors"
)



type MovieService struct{}

func NewMovieService() *MovieService{
	return &MovieService{}
}

func(s *MovieService) List() []models.Movie {
	return repo.GetAll()
}

func (s *MovieService) Get(id string) (*models.Movie, error){
	m := repo.GetById(id)
	if m == nil {
		return nil, errors.New("Movie Not Found")
	}
	return m, nil
}

func (s *MovieService) Create(m models.Movie) (models.Movie, error){
	if m.Title == "" || m.Director == nil {
		return models.Movie{}, errors.New("Title and Director Required")
	}
	created := repo.CreateMovie(m)
	return created, nil
}

func (s *MovieService) Update(id string, m models.Movie) (*models.Movie, error){
	updated := repo.UpdateMovie(id, m)
	if updated  == nil{
		return  nil, errors.New("Movie not Found")
	}
	return  updated, nil
}

func (s *MovieService) Delete(id string) ([]models.Movie, error){
	remaining := repo.DeleteMovie(id)
	return remaining, nil
}