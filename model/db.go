package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "golaxy_admin"
	password = "supermassive_black_hole"
	dbname   = "golaxy"
)

var Db *sql.DB

func main() {
	dbCredentials := fmt.Sprintf("host=%s port=%d user=%s" +
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err := sql.Open("postgres", dbCredentials)

	if errors.Is(err, nil) {
		panic(err)
	}

	defer Db.Close()
}
