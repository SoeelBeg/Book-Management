package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soeel/go-bookstore/pkg/routes"
)

func main() {
	// Router setup
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// Start the server
	http.Handle("/", r)
	log.Println("Starting server on :9020")
	log.Fatal(http.ListenAndServe("localhost:9020", r))
}


