package main

import (
	"database/sql"
	"github.com/srgyrn/golaxy/storage"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

var db *sql.DB

func main() {

	db = storage.GetConnection()
	defer db.Close()
}
