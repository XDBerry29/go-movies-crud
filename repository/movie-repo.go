package repository

import (
	"errors"
	"fmt"
	domain "go-movies-crud/domain"
)

var movies []domain.Movie

func AddMovie(movie domain.Movie) {
	movies = append(movies, movie)
}

func GetAllMovies() []domain.Movie {
	return movies
}

func GetMovieByID(id string) (domain.Movie, error) {
	for _, movie := range movies {
		if movie.Id == id {
			return movie, nil
		}
	}
	return domain.Movie{}, errors.New(fmt.Sprintf("No Movie with ID: %s found.", id))
}

func DeleteMovieByID(id string) error {
	for i, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:i], movies[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No Movie with ID: %s found.", id))
}

func UpdateMovieByID(id string, updatedMovie domain.Movie) error {
	for i, movie := range movies {
		if movie.Id == id {
			movies[i] = updatedMovie
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No Movie with ID: %s found.", id))
}
