package director

import (
	"database/sql"
	"encoding/json"
	"errors"
	directorResponse "github.com/srgyrn/golaxy/controller"
	"github.com/srgyrn/golaxy/storage"
	"io/ioutil"
	"net/http"
	"strconv"
)

var directorGW *storage.DirectorGateway

func init() {
	directorGW = &storage.DirectorGateway{}
}

func MakeHandlerFunction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createDirector(w, r)
		case "GET":
			//TODO: Find a better solution
			id, _ := strconv.Atoi(r.URL.Path[len("/directors/"):])
			if 0 == id {
				returnAllDirector(w, r)
				return
			}

			returnDirectorByID(w, r)
		case "PUT":
			updateDirectorByID(w, r)
		case "DELETE":
			deleteDirectorByID(w, r)
		default:
			http.NotFound(w, r)
			return
		}
	}
}

func createDirector(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var director storage.Director
	err := json.Unmarshal(body, &director)

	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = directorGW.Insert(&director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	directorResponse.SuccessfulResponseListener(w, &directorResponse.Response{Data: &director})
	return
}

func returnAllDirector(w http.ResponseWriter, r *http.Request) {
	directors, err := directorGW.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	directorResponse.SuccessfulResponseListener(w, &directorResponse.Response{Data: &directors})
}

func returnDirectorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/directors/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	director, err := directorGW.Find(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "director not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	directorResponse.SuccessfulResponseListener(w, &directorResponse.Response{Data: &director})
}

func updateDirectorByID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var director storage.Director

	err := json.Unmarshal(body, &director)
	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = directorGW.Update(&director)

	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	directorResponse.SuccessfulResponseListener(w, &directorResponse.Response{Data: &director})
}

func deleteDirectorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/directors/"):])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = directorGW.Delete(id)
	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	directorResponse.SuccessfulResponseListener(w, &directorResponse.Response{Data: "director deleted successfully"})
}
