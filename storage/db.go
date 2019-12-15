package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "golaxy_admin"
	password = "supermassive_black_hole"
	dbname   = "golaxy"
)

// TableGateway defines the functions every table gateway should implement in the project.
type TableGateway interface {
	Find(id int) error
	FindAll() error
	Insert() error
	Update() (int64, error)
	Delete(id int) (int64, error)
}

var db *sql.DB

// InitDB initiates the database connection
func InitDB() {
	var err error
	dbCredentials := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", dbCredentials)

	if !errors.Is(err, nil) {
		log.Panic(err)
	}

	err = db.Ping()

	/*if !errors.Is(err, nil) {
		log.Panic(err)
	}*/

	fmt.Println("Database connection successful!")
}

// CloseDB closes the database connection.
func CloseDB() {
	err := db.Close()
	if !errors.Is(err, nil) {
		log.Fatal(err.Error())
	}
}
