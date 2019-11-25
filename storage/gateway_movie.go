package storage

import (
	"github.com/lib/pq"
	"log"
)

// MovieGateway is a type of TableGateway
type MovieGateway struct{}

// Find retrieves the movie by id
func (gw MovieGateway) Find(id int) (*Movie, error) {
	movie := new(Movie)
	err := db.QueryRow(`
		SELECT m.id,
		       m.name,
		       m.year,
		       genre,
		       rating,
		       movie_cast,
		       m.created_at,
		       d.id as director_id,
		       d.name as director_name,
		       d.surname as director_surname,
		       d.birth_year,
		       d.created_at as director_created_at
		FROM movies as m
		INNER JOIN directors d on m.director_id = d.id
		WHERE m.id = $1
		`, id).Scan(&movie.ID,
			&movie.Name,
			&movie.Year,
			pq.Array(&movie.Genre),
			&movie.Rating,
			pq.Array(&movie.Cast),
			&movie.CreatedAt,
			&movie.Director.ID,
			&movie.Director.Name,
			&movie.Director.Surname,
			&movie.Director.Year,
			&movie.Director.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return movie, nil
}

// FindAll retrieves all the movies
func (gw MovieGateway) FindAll() ([]Movie, error) {
	rows, err := db.Query(`
		SELECT m.id,
		       m.name,
		       m.year,
		       genre,
		       rating,
		       movie_cast,
		       m.created_at,
		       d.id as director_id,
		       d.name as director_name,
		       d.surname as director_surname,
		       d.birth_year,
		       d.created_at as director_created_at
		FROM movies as m
		INNER JOIN directors d on m.director_id = d.id
		`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []Movie

	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.ID,
			&movie.Name,
			&movie.Year,
			pq.Array(&movie.Genre),
			&movie.Rating,
			pq.Array(&movie.Cast),
			&movie.CreatedAt,
			&movie.Director.ID,
			&movie.Director.Name,
			&movie.Director.Surname,
			&movie.Director.Year,
			&movie.Director.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		movies = append(movies, *movie)
	}

	return movies, nil
}

// Insert inserts a new row to the movies table
func (gw MovieGateway) Insert(movie *Movie) (*Movie, error) {
	err := db.QueryRow("INSERT INTO movies (name, rating, year, movie_cast, director_id, genre) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		movie.Name,
		movie.Rating,
		movie.Year,
		pq.Array(movie.Cast),
		movie.Director.ID,
		pq.Array(movie.Genre),
	).Scan(&movie.ID)

	if err != nil {
		log.Printf("error at creating movie. Data: %v", movie)
		return nil, err
	}

	//TODO: return movie with director data
	return movie, nil
}

// Update updates a row in movies table by id
func (gw MovieGateway) Update(movie *Movie) (int64, error) {
	result, err := db.Exec("UPDATE movies SET name = $2, rating = $3, year = $4, movie_cast = $5, director_id = $6, genre = $7 WHERE id = $1",
		movie.ID,
		movie.Name,
		movie.Rating,
		movie.Year,
		pq.Array(movie.Cast),
		movie.Director.ID,
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
