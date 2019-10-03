package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/srgyrn/golaxy/storage"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

var db *sql.DB

type MoviePostHandler struct{}

func (mph MoviePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	movieGw := storage.NewMovieGateway(db)
	movie := storage.Movie{
		Name:     "Batman",
		Year:     1989,
		Director: "Tim Burton",
		Rating:   7.5,
		Genre:    []string{"action", "adventure"},
		Cast:     []string{"Michael Keaton", "Jack Nicholson", "Kim Basinger"},
	}

	_, err := movieGw.Insert(&movie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, movie.Id)
}

type MovieGetHandler struct{}

func (mph MovieGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// id, err := strconv.Atoi(r.URL.Path[len("/movies/"):])
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// movieGW := model.NewMovieGateway(db)
	// movie, err := movieGW.Find(id)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		http.Error(w, "", http.StatusNotFound)
	// 		return
	// 	}

	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// writer.Header().Set("Content-Type", "application/json")
	// writer.WriteHeader(200)
	// if err = json.NewEncoder(writer).Encode(&movie); err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

type MovieHandler struct{}

func (mph MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		mph := MoviePostHandler{}
		return mph.ServeHTTP(w, r)
	case "GET":
		mgh := MovieGetHandler{}
		return mgh.ServeHTTP(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func main() {

	db = storage.GetConnection()
	defer db.Close()

	http.Handle("/movies/", MovieHandler{})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
