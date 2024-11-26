package handlers

import (
	"api/model"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	movies []model.Movie
	mu     sync.Mutex
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if params["id"] == movie.ID {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
	mu.Lock()
	movies = append(movies, movie)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedMovie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies[i] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, movie := range movies {
		if params["id"] == movie.ID {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}
