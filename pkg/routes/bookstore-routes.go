package routes

import (
	"github.com/gorilla/mux"
	"github.com/heloise-viegas/Go-BookStore/pkg/controllers"
)

var RegisteredBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	roter.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
