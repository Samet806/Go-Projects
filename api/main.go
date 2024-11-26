package main

import (
	"api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovieByID).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
