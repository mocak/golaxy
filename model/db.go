package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "golaxy_admin"
	password = "supermassive_black_hole"
	dbname   = "golaxy"
)

type TableGateway interface {
	Find(id int) (*struct{}, error)
	FindAll() (*struct{}, error)
	Insert(*struct{}) (*struct{}, error)
	Update(*struct{}) (int64, error)
	Delete(*struct{}) (int64, error)
}

func getConnection() *sql.DB {
	dbCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbCredentials)

	if !errors.Is(err, nil) {
		log.Panic(err)
	}

	err = db.Ping()

	if !errors.Is(err, nil) {
		log.Panic(err)
	}

	fmt.Println("Database connection successful!")

	return db
}

func BootstrapDatabase() {
	db := getConnection()
	defer db.Close()

	createMovieQuery := `
		DROP TABLE IF EXISTS movies;
		CREATE TABLE movies (
			id serial PRIMARY KEY, 
			name varchar (200) NOT NULL, 
			year int NULL CHECK (year > 0), 
			genre text [] NULL, 
			rating numeric (2), 
			director varchar (50), 
			movie_cast text [], 
			created_at timestamp DEFAULT CURRENT_TIMESTAMP, 
			UNIQUE (name, year)
			)`

	stmt, err := db.Prepare(createMovieQuery)
	checkError("preparing statement", err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkError("executing statement", err)

	log.Println("Table creation end.")
}

func checkError(stepDescription string, err error) {
	if errors.Is(err, nil) {
		return
	}

	log.Printf("Table creation failed at %s | %s\n", stepDescription, err.Error())
}
