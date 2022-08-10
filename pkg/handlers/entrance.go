package handlers

import (
	"cofeeShop/storage"
	"log"
	"net/http"
	"os"
)

func Enter(w http.ResponseWriter, r *http.Request) {

	enterFile, err := os.ReadFile(storage.GetFileConfig().CommonInfoFile)
	if err != nil {
		http.Error(w, "data is not available", 204)
		log.Fatal(err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(enterFile)
	if err != nil {
		http.Error(w, "sending error", 500)
		log.Fatal(err.Error())
	}

}
