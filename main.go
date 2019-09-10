package main

import (
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
	model.CreateTableMovie()
}
