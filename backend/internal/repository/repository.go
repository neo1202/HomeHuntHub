package repository //因為在repository folder中

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	//list the function this repo must have
	Connection() *sql.DB                 //return DB pointer
	AllMovies() ([]*models.Movie, error) //takes no params, return a slice
}
