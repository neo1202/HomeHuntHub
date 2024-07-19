package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

// 這個struct會符合repository.go中的interface因為他implement了interface
type PostgresDBRepo struct {
	//hold our connection to DB
	//(a pointer
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// 以下皆為實現interface的function在這個struct, 前方(m *PostgresDBRepo)是指receiver

// 單純獲取他的connection(conn), 不是在此處連結
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

// takes no params, return a pointer to models.movie with potentially an error
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//golang很討厭null value
	//coalesce可以讓如果沒找到返回空字串
	query := `
		select
			id,title,release_date,runtime,
			mpaa_rating,description,coalesce(image, ''),
			created_at, updated_at
		from
			movies
		order by
			title
	`
	rows, err := m.DB.QueryContext(ctx, query) //你有ctx秒去做這個query
	if err != nil {
		return nil, err
	}
	defer rows.Close() //壹定要關掉db不然結束後他還是開著會太久

	var movies []*models.Movie

	for rows.Next() { //scan可以拿來一行行把db讀取到的寫進golang的結構
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
					created_at, updated_at from users where email = $1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, password,
					created_at, updated_at from users where id = $1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
