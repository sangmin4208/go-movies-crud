package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getMovies")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteMovie")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for i, item := range movies {
		if item.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getMovie")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createMovie")
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updateMovie")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for i, item := range movies {
		if item.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			movie.ID = id
			movies = append(movies, movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Title: "The Shawshank Redemption", Director: &Director{FirstName: "Frank", LastName: "Darabont"}})
	movies = append(movies, Movie{ID: "2", Title: "The Godfather", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "3", Title: "The Godfather: Part II", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "4", Title: "The Dark Knight", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "5", Title: "12 Angry", Director: &Director{FirstName: "Ridley", LastName: "Scott"}})
	movies = append(movies, Movie{ID: "6", Title: "Schindler's List", Director: &Director{FirstName: "Steven", LastName: "Spielberg"}})
	movies = append(movies, Movie{ID: "7", Title: "The Lord of the Rings: The Return of the King", Director: &Director{FirstName: "Peter", LastName: "Jackson"}})
	movies = append(movies, Movie{ID: "8", Title: "Pulp Fiction", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "9", Title: "Fight Club", Director: &Director{FirstName: "David", LastName: "Fincher"}})
	movies = append(movies, Movie{ID: "10", Title: "Inception", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
