package service

import (
	domain "go-movies-crud/domain"
	repository "go-movies-crud/repository"
)

func AddMovie(newMovie domain.Movie) {
	repository.AddMovie(newMovie)
}

func GetAllMovies() []domain.Movie {
	return repository.GetAllMovies()
}

func DeleteMovieByID(id string) error {
	return repository.DeleteMovieByID(id)
}

func GetMovieByID(id string) (domain.Movie, error) {
	return repository.GetMovieByID(id)
}

func UpdateMovieByID(id string, updated domain.Movie) error {
	return repository.UpdateMovieByID(id, updated)
}
