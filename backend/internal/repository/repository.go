package repository //因為在repository folder中

import "backend/internal/models"

type DatabaseRepo interface {
	//list the function this repo must have
	AllMovies()([]*models.Movie, error)//takes no params, return a slice
}