package routes

import (
	"github.com/gorilla/mux"
	"github.com/soeel/go-bookstore/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId:[0-9]+}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId:[0-9]+}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId:[0-9]+}", controller.DeleteBook).Methods("DELETE")
}
