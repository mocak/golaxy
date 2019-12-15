package controller

import (
	"encoding/json"
	"errors"
	"github.com/srgyrn/golaxy/storage"
	"io/ioutil"
	"net/http"
)

type MovieHandler struct {
	path string
	gateway *storage.MovieGateway
}

func newHandler() *CommonHandler {
	return &CommonHandler{
		path: "/movies",
		gw: storage.MovieGateway{},
	}
}

func (mh *MovieHandler) create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var movie storage.Movie
	err := json.Unmarshal(body, &movie)

	if !errors.Is(err, nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = mh.gateway.Insert(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successfulResponseListener(w, &Response{Data: &movie})
	return
}

func (mh *MovieHandler) getAll(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.gateway.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	successfulResponseListener(w, &Response{Data: &movies})
}