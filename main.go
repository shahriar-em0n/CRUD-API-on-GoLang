package main

import (
	// "crypto/rand"
	// "math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "github.com/gogo/protobuf/test/indeximport-issue72/index"
	"github.com/gorilla/mux"
// "	go get google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
)

type Movie struct {
	ID       string    `json:"id"`
	IsBn     string    `json:"isbn"`
	Titile   string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	maxID := 0
	for _, m := range movies {
		id, _ := strconv.Atoi(m.ID)
		if id > maxID {
			maxID = id
		}
	}
	movie.ID = strconv.Itoa(maxID + 1)

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// func createMovie(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type","application/json")
// 	var movie Movie
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID = strconv.Itoa(rand.Intn(100000))
// 	movies = append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
// }

func updateMovie(w http.ResponseWriter, r *http.Request){
	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: "1", IsBn: "438227", Titile: "Movie One", Director: &Director{FirstName: "Shahriar", LastName: "Stranger"}},
		Movie{ID: "2", IsBn: "739012", Titile: "Echoes of Time", Director: &Director{FirstName: "Hakim", LastName: "Akash"}},
		Movie{ID: "3", IsBn: "562934", Titile: "Quantum Mirage", Director: &Director{FirstName: "Liam", LastName: "Noir"}},
		Movie{ID: "4", IsBn: "823675", Titile: "Fragments of Tomorrow", Director: &Director{FirstName: "Ava", LastName: "Lumen"}},
		Movie{ID: "5", IsBn: "982341", Titile: "Silent Horizon", Director: &Director{FirstName: "Noah", LastName: "Vale"}},
		Movie{ID: "6", IsBn: "114589", Titile: "The Last Algorithm", Director: &Director{FirstName: "Elena", LastName: "Cross"}},
		Movie{ID: "7", IsBn: "7781178", Titile: "Quats Lots", Director: &Director{FirstName: "Jhon", LastName: "Deh"}},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
