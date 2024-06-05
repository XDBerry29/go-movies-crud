package controller

import (
	"encoding/json"
	domain "go-movies-crud/domain"
	service "go-movies-crud/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetAllMovies())
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := service.DeleteMovieByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie, err := service.GetMovieByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(movie)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie domain.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	service.AddMovie(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie domain.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	params := mux.Vars(r)
	err := service.UpdateMovieByID(params["id"], movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
