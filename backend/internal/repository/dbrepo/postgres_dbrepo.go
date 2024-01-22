package dbrepo

import (
	"backend/internal/models"
	"database/sql"
	"time"
)

//這個struct會符合repository.go中的interface因為他implement了interface
type PostgresDBRepo struct {
	//hold our connection to DB
	//(a pointer
		DB *sql.DB
}

const dbTimeout = time.Second*3

// takes no params, return a pointer to models.movie with potentially an error
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error){
	var movies []*models.Movie
	return movies,nil
}