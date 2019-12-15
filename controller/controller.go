package controller

import (
	"net/http"
	"strconv"

	"github.com/srgyrn/golaxy/storage"
)

type GolaxyHandler interface {
	NewHandler() GolaxyHandler
	create()
	getByID()
	getAll()
	deleteByID()
	updateByID()
}

type CommonHandler struct {
	handler *GolaxyHandler
	path string
	gw storage.TableGateway
}


func RegisterController(c *CommonHandler) {

	http.Handle(c.path, c.handler)
}


func (h *CommonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.
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