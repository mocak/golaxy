package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/srgyrn/golaxy/model"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

var db *sql.DB

func main() {

	db = model.GetConnection()
	defer db.Close()

	//model.BootstrapDatabase(db)

	http.HandleFunc("/movies/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			id, err := strconv.Atoi(request.URL.Path[len("/movies/"):])
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}

			movieGW := model.NewMovieGateway(db)
			movie, err := movieGW.Find(id)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(writer, "", http.StatusNotFound)
					return
				}

				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(200)
			if err = json.NewEncoder(writer).Encode(&movie); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(writer, "", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func insertMovie(movieGw *model.MovieGateway) {

	movie := model.Movie{
		Name:     "Batman",
		Year:     1989,
		Director: "Tim Burton",
		Rating:   7.5,
		Genre:    []string{"action", "adventure"},
		Cast:     []string{"Michael Keaton", "Jack Nicholson", "Kim Basinger"},
	}

	_, err := movieGw.Insert(&movie)

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(movie)
}

func getMovie(movieGw *model.MovieGateway) {
	movie, err := movieGw.Find(1)

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(movie)
}
