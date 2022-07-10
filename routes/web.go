package routes

import (
	"log"
	"net/http"

	IndexController "github.com/SojebSikder/goframe/app/controller"
	"github.com/julienschmidt/httprouter"
)

func Routes(router *httprouter.Router) {
	router.GET("/", IndexController.Index)
	router.GET("/hello/:name", IndexController.Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
