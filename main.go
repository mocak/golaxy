package main

import (
	"database/sql"
	"github.com/srgyrn/golaxy/storage"
	"log"
	"net/http"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

var db *sql.DB

type MoviePostHandler struct {}


type MovieHandler struct {}

func (mph MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		return
	}
}

func main() {

	db = storage.GetConnection()
	defer db.Close()

	http.Handle("/movies/", MovieHandler{})




	log.Fatal(http.ListenAndServe(":8090", nil))
}
