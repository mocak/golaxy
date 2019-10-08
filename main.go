package main

import (
	"github.com/srgyrn/golaxy/controller/director"
	"github.com/srgyrn/golaxy/controller/movie"
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

func main() {
	storage.InitDB()
	defer storage.CloseDB()

	handleRequests()
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func handleRequests() {
	http.HandleFunc("/movies/", movie.MakeHandlerFunction())
	http.HandleFunc("/directors/", director.MakeHandlerFunction())
}
