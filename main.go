package main

import (
	"log"
	"net/http"
	"go_api/router"
)

func main() {
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8090", router))
}
