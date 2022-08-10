package handlers

import (
	"cofeeShop/pkg/entities"
	"cofeeShop/storage"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func ShowMenu(w http.ResponseWriter, r *http.Request) {

	enterFile, err := os.ReadFile(storage.GetFileConfig().CommonInfoFile)
	if !isShopOpen(enterFile) {
		_, err = w.Write([]byte("Закрыто"))
		return
	}

	menuFile, err := os.ReadFile(storage.GetFileConfig().MenuFile)
	if err != nil {
		http.Error(w, "data is not available", 204)
		log.Fatal(err.Error())
	}

	_, err = w.Write(menuFile)
	if err != nil {
		http.Error(w, "sending error", 500)
		log.Fatal(err.Error())
	}

}

func isShopOpen(encodedGeneralInfo []byte) bool {
	info := entities.ShopInfo{}

	err := json.Unmarshal(encodedGeneralInfo, &info)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	if info.IsOpen {
		return true
	}

	return false
}
