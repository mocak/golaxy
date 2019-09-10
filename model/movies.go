package model

type Movie struct {
	Id        int
	Name      string
	Year      int
	Genre     []string
	Rating    float32
	Director  string
	Cast      []string
	CreatedAt string
}

func CreateTableMovie() {
	createTable("movie",
		`CREATE TABLE movie (
			id serial PRIMARY KEY, 
			name varchar (200) NOT NULL, 
			year int NULL CHECK (year > 0), 
			genre text [] NULL, 
			rating numeric (2), 
			director varchar (50), 
			movie_cast text [], 
			created_at timestamp DEFAULT CURRENT_TIMESTAMP, 
			UNIQUE (name, year)
			)`)
}
