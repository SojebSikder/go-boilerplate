package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/SojebSikder/goframe/routes"
)

func main() {
	router := httprouter.New()

	routes.Routes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
