package main

import (
	"github.com/heloise-viegas/Go-BookStore/pkg/routes"
	"github.com/jinzhu/gorm/dialects/mysql"
	"hithub.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
