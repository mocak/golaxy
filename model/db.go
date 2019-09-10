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

func initDB() *sql.DB {
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

func createTable(tableName, createScript string) {
	db := initDB()
	defer db.Close()

	stmt, err := db.Prepare(createScript)
	checkError("preparing statement", tableName, err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkError("executing statement", tableName, err)

	log.Println("Table creation end.")
}

func checkError(stepDescription, tableName string, err error) {
	if errors.Is(err, nil) {
		return
	}

	log.Printf("Table creation failed at %s | table: %s | %s\n", stepDescription, tableName, err.Error())
}
