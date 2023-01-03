package router

import (
	"go-learn/controller/cake"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/cakes", cake.Create).Methods("POST")
	router.HandleFunc("/cakes", cake.GetAll).Methods("GET")
	router.HandleFunc("/cakes/{id}", cake.Update).Methods("PUT")
	router.HandleFunc("/cakes/{id}", cake.GetOne).Methods("GET")
	router.HandleFunc("/cakes/{id}", cake.Delete).Methods("DELETE")
	return router
}
