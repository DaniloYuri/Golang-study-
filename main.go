package main

import (
	"Golang/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
