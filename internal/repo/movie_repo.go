package repo

import (
	"CRUD-API/internal/models"
	"strconv"

	// "github.com/go-playground/validator/v10/translations/id"
)

var movies []models.Movie

// Seed initial movies (same data you provided)
func Seed() {
	movies = []models.Movie{
		{ID: "1", IsBn: "438227", Title: "Movie One", Director: &models.Director{FirstName: "Shahriar", LastName: "Stranger"}},
		{ID: "2", IsBn: "739012", Title: "Echoes of Time", Director: &models.Director{FirstName: "Hakim", LastName: "Akash"}},
		{ID: "3", IsBn: "562934", Title: "Quantum Mirage", Director: &models.Director{FirstName: "Liam", LastName: "Noir"}},
		{ID: "4", IsBn: "823675", Title: "Fragments of Tomorrow", Director: &models.Director{FirstName: "Ava", LastName: "Lumen"}},
		{ID: "5", IsBn: "982341", Title: "Silent Horizon", Director: &models.Director{FirstName: "Noah", LastName: "Vale"}},
		{ID: "6", IsBn: "114589", Title: "The Last Algorithm", Director: &models.Director{FirstName: "Elena", LastName: "Cross"}},
		{ID: "7", IsBn: "7781178", Title: "Quats Lots", Director: &models.Director{FirstName: "Jhon", LastName: "Deh"}},
	}
}

func GetAll() []models.Movie{
	return movies
}

func GetById(id string) *models.Movie{
	for _, m:= range movies{
		if m.ID == id {
			// return a pointer to a copy stored in slice
			mCopy := m
			return &mCopy
		}
	}
	return nil
}

func CreateMovie(moive models.Movie) models.Movie{
	maxId := 0
	for _, m := range movies{
		id, err := strconv.Atoi(m.ID)
		if err != nil {
			continue
		}
		if id > maxId {
			maxId = id
		}
	}
	moive.ID = strconv.Itoa(maxId + 1)
	movies = append(movies, moive)
	return moive
}

func UpdateMovie(id string, update models.Movie) *models.Movie{
	for i, m := range movies{
		if m.ID == id{
			update.ID = id
			movies[i] = update
			mCopy := movies[i]
			return &mCopy
		}
	}
	return nil
}

func DeleteMovie(id string) []models.Movie{
	for i, m := range movies{
		if m.ID == id {
			movies = append(movies[:i], movies[i + 1:]...)
			break
		}
	}
	return movies
}