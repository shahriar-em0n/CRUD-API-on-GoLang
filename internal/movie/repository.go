package movie

import "strconv"

var movies []Movie

func SeedData() {
	movies = []Movie{
		{ID: "1", IsBn: "438227", Title: "Movie One", Director: &Director{FirstName: "Shahriar", LastName: "Stranger"}},
		{ID: "2", IsBn: "739012", Title: "Echoes of Time", Director: &Director{FirstName: "Hakim", LastName: "Akash"}},
		{ID: "3", IsBn: "562934", Title: "Quantum Mirage", Director: &Director{FirstName: "Liam", LastName: "Noir"}},
		{ID: "4", IsBn: "823675", Title: "Fragments of Tomorrow", Director: &Director{FirstName: "Ava", LastName: "Lumen"}},
		{ID: "5", IsBn: "982341", Title: "Silent Horizon", Director: &Director{FirstName: "Noah", LastName: "Vale"}},
		{ID: "6", IsBn: "114589", Title: "The Last Algorithm", Director: &Director{FirstName: "Elena", LastName: "Cross"}},
		{ID: "7", IsBn: "7781178", Title: "Quats Lots", Director: &Director{FirstName: "Jhon", LastName: "Deh"}},
	}
}

func GetAll() []Movie {
	return movies
}

func GetById(id string) *Movie {
	for _, m := range movies {
		if m.ID == id {
			return &m
		}
	}
	return nil
}

func CreateMovie(movie Movie) Movie {
	maxId := 0
	for _, m := range movies {
		id, err := strconv.Atoi(m.ID)
		if err != nil {
			continue
		}

		if id > maxId {
			maxId = id
		}
	}
	movie.ID = strconv.Itoa(maxId + 1)
	movies = append(movies, movie)
	return movie
}

func updateMovie(id string, update Movie) *Movie{
	for i, m := range movies{
		if m.ID == id {
			update.ID = id
			movies[i] = update
			return &movies[i]
		}
	}
	return nil
}

func deleteMovie(id string) []Movie{
	for i, m := range movies{
		if m.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	return movies
}