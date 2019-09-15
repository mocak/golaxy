package main

import (
	"fmt"
	"github.com/srgyrn/golaxy/model"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	//model.BootstrapDatabase()

	movieGw := new(model.MovieGateway)
	//insertMovie(movieGw)
	getMovie(movieGw)
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