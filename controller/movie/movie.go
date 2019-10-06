package movie

import (
	"database/sql"
	"encoding/json"
	"errors"
	movieResponse "github.com/srgyrn/golaxy/controller"
	"github.com/srgyrn/golaxy/storage"
	"io/ioutil"
	"net/http"
	"strconv"
)

var movieGW *storage.MovieGateway

func init() {
	movieGW = &storage.MovieGateway{}
}

func MakeHandlerFunction() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createMovie(w, r)
		case "GET":
			//TODO: Find a better solution
			id, _ := strconv.Atoi(r.URL.Path[len("/movies/"):])
			if 0 == id {
				returnAllMovie(w, r)
				return
			}

			returnMovieByID(w, r)
		case "PUT":
			updateMovieByID(w,r)
		case "DELETE":
			deleteMovieByID(w, r)
		default:
			http.NotFound(w, r)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var movie storage.Movie
	err := json.Unmarshal(body, &movie)

	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = movieGW.Insert(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieResponse.SuccessfulResponseListener(w, &movieResponse.Response{Data: &movie})
	return
}

func returnAllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := movieGW.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieResponse.SuccessfulResponseListener(w, &movieResponse.Response{Data: &movies})
}

func returnMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/movies/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie, err := movieGW.Find(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "movie not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieResponse.SuccessfulResponseListener(w, &movieResponse.Response{Data: &movie})
}

func updateMovieByID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var movie storage.Movie

	err := json.Unmarshal(body, &movie)
	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = movieGW.Update(&movie)

	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieResponse.SuccessfulResponseListener(w, &movieResponse.Response{Data: &movie})
}

func deleteMovieByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/movies/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = movieGW.Delete(id)
	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	movieResponse.SuccessfulResponseListener(w, &movieResponse.Response{Data: "movie deleted successfully"})
}
