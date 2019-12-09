package controller

import (
	"net/http"
	"strconv"
)

type CommonHandler struct {

}

func (m CommonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		create(w, r)
	case http.MethodGet:
		//TODO: Find a better solution
		id, _ := strconv.Atoi(r.URL.Path[len("/movies/"):])
		if 0 == id {
			getAll(w, r)
			return
		}
		getByID(w, r)
	case http.MethodPut:
		updateByID(w, r)
	case http.MethodDelete:
		deleteByID(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}

type MovieHandler struct {
	h CommonHandler
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}