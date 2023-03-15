package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/awilson506/mlb-stats-api/api"
	"github.com/awilson506/mlb-stats-api/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":8080", "change port to run service on.  defaults 8080.")
}

func main() {
	flag.Parse()

	// start the server
	s := server.New(port, api.New())
	err := s.Start()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
