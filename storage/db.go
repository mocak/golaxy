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

type TableGateway interface {
	Find(id int) (*struct{}, error)
	FindAll() (*struct{}, error)
	Insert(*struct{}) (*struct{}, error)
	Update(*struct{}) (int64, error)
	Delete(*struct{}) (int64, error)
}

func GetConnection() *sql.DB {
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
