package app

import (
	"cofeeShop/pkg/config"
	"cofeeShop/pkg/server"
	"cofeeShop/storage"
	"github.com/gorilla/mux"
)

func Run() {
	var serverObject = server.Server{
		Config: config.Config{Address: "127.0.0.1:8080"},
		Mux:    mux.NewRouter(),
		FileConfig: storage.FileConfig{
			CommonInfoFile:   "storage/json/general_info.json",
			MenuFile:         "storage/json/menu",
			ProductsListFile: "storage/json/products",
		},
	}

	serverObject.Start()
}
