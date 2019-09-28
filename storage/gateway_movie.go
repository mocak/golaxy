package storage

import (
	"database/sql"

	"github.com/lib/pq"
)

type MovieGateway struct {
	db *sql.DB
}

func NewMovieGateway(db *sql.DB) *MovieGateway {
	return &MovieGateway{db: db}
}

func (gw MovieGateway) Find(id int) (*Movie, error) {

	movie := new(Movie)
	err := gw.db.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(&movie.Id, &movie.Name, &movie.Year, pq.Array(&movie.Genre), &movie.Rating, &movie.Director, pq.Array(&movie.Cast), &movie.CreatedAt)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (gw MovieGateway) FindAll() ([]*Movie, error) {

	rows, err := gw.db.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []*Movie

	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, pq.Array(&movie.Genre), &movie.Rating, &movie.Director, pq.Array(&movie.Cast), &movie.CreatedAt)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (gw MovieGateway) Insert(movie *Movie) (*Movie, error) {

	err := gw.db.QueryRow("INSERT INTO movies (name, rating, year, movie_cast, director, genre) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		movie.Name,
		movie.Rating,
		movie.Year,
		pq.Array(movie.Cast),
		movie.Director,
		pq.Array(movie.Genre),
	).Scan(&movie.Id)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (gw MovieGateway) Update(movie *Movie) (int64, error) {

	result, err := gw.db.Exec("UPDATE movies SET name = $2, rating = $3, year = $4, movie_cast = $5, director = $6, genre = $7 WHERE id = $1",
		movie.Id,
		movie.Name,
		movie.Rating,
		movie.Year,
		pq.Array(movie.Cast),
		movie.Director,
		pq.Array(movie.Genre),
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (gw MovieGateway) Delete(id int) (int64, error) {

	result, err := gw.db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
