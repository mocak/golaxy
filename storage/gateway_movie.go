package storage

import (
	"github.com/lib/pq"
)

// MovieGateway is a type of TableGateway
type MovieGateway struct{}

// Find retrieves the movie by id
func (gw MovieGateway) Find(id int) (*Movie, error) {
	movie := new(Movie)
	err := db.QueryRow("SELECT * FROM movies WHERE id = $1", id).Scan(&movie.ID, &movie.Name, &movie.Year, pq.Array(&movie.Genre), &movie.Rating, &movie.Director, pq.Array(&movie.Cast), &movie.CreatedAt)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

// FindAll retrieves all the movies
func (gw MovieGateway) FindAll() ([]*Movie, error) {
	rows, err := db.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []*Movie

	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Year, pq.Array(&movie.Genre), &movie.Rating, &movie.Director, pq.Array(&movie.Cast), &movie.CreatedAt)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// Insert inserts a new row to the movies table
func (gw MovieGateway) Insert(movie *Movie) (*Movie, error) {
	err := db.QueryRow("INSERT INTO movies (name, rating, year, movie_cast, director, genre) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		movie.Name,
		movie.Rating,
		movie.Year,
		pq.Array(movie.Cast),
		movie.Director,
		pq.Array(movie.Genre),
	).Scan(&movie.ID)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

// Update updates a row in movies table by id
func (gw MovieGateway) Update(movie *Movie) (int64, error) {
	result, err := db.Exec("UPDATE movies SET name = $2, rating = $3, year = $4, movie_cast = $5, director = $6, genre = $7 WHERE id = $1",
		movie.ID,
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

// Delete removes a row in table movies by id
func (gw MovieGateway) Delete(id int) (int64, error) {
	result, err := db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
