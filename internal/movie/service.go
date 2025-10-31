package movie

import "errors"


type MovieService struct{}

func NewMovieService() *MovieService{
	return &MovieService{}
}

func (s *MovieService) List() []Movie{
	return GetAll()
}

func (s *MovieService) Get(id string) (*Movie, error){
	m := GetById(id)
	if m == nil{
		return nil, errors.New("Movie not found")
	}
	return m, nil
}

func (s *MovieService) Create(m Movie)(Movie, error){
	if m.Title == "" || m.Director == nil {
		return Movie{}, errors.New("Title and Director Required")
	}
	created:= CreateMovie(m)
	return created, nil
}

func (s *MovieService) Update(id string, m Movie) (*Movie, error){
	update := updateMovie(id, m)
	if update == nil {
		return nil, errors.New("Movie Not Found")
	}
	return update, nil
}

func (s *MovieService) Delete(id string) ([]Movie, error){
	return deleteMovie(id),nil
}