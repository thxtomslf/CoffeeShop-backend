package routes

import (
	"cofeeShop/pkg/handlers"
	"github.com/gorilla/mux"
)

func ConfigRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.Enter).Methods("GET")

	router.HandleFunc("/menu", handlers.ShowMenu).Methods("GET")

	router.HandleFunc("/cook", handlers.CookProducts).Methods("POST")
}
