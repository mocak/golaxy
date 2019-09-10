package model

import (
	"database/sql"
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

func init() {
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s" +
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
