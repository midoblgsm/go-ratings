package main

import (
	"log"

	"github.com/midoblgsm/go-ratings/web_server"
)

func main() {

	server := web_server.NewRatingsApiServer(9998)

	log.Fatal(server.Start())
}
