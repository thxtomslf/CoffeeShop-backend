package server

import (
	"cofeeShop/pkg/config"
	"cofeeShop/pkg/routes"
	"cofeeShop/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	config.Config
	storage.FileConfig
	Mux *mux.Router
}

func (server *Server) Start() {
	routes.ConfigRoutes(server.Mux)

	err := http.ListenAndServe(server.Config.Address, server.Mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
