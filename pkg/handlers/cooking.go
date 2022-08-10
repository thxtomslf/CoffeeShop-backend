package handlers

import (
	"cofeeShop/pkg/entities"
	"cofeeShop/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CookProducts(w http.ResponseWriter, r *http.Request) {
	var cookingData []byte
	cookingData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	coffeeToCook := string(cookingData)
	fmt.Println(coffeeToCook)

	productsData, err := os.ReadFile(storage.GetFileConfig().ProductsListFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	commonProductList := entities.CommonProductList{}
	err = json.Unmarshal(productsData, &commonProductList)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(commonProductList)

	if products, ok := isEnoughProducts(coffeeToCook, &commonProductList); ok {
		cook(products, &commonProductList)

		updateProductList(&commonProductList)

		_, err = w.Write([]byte(`{"done: true"}`))
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		_, err := w.Write([]byte(`{"done: false"}`))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

func updateProductList(commonProductList *entities.CommonProductList) {
	updatedProductList, err := json.Marshal(commonProductList)
	if err != nil {
		log.Fatal(err.Error())
	}

	productFile, err := os.OpenFile(storage.GetFileConfig().ProductsListFile, os.O_WRONLY|os.O_TRUNC, 0777)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer productFile.Close()

	_, err = productFile.Write(updatedProductList)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func isEnoughProducts(coffeeToCook string, commonList *entities.CommonProductList) ([]string, bool) {
	menuData, err := os.ReadFile(storage.GetFileConfig().MenuFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	menu := entities.Menu{}
	err = json.Unmarshal(menuData, &menu)
	if err != nil {
		log.Fatal(err.Error())
	}

	var products []string
	for _, coffeeList := range menu {
		for coffee, info := range coffeeList {
			if coffee == coffeeToCook {
				products = info.Composition
				return checkAvailability(products, commonList)
			}
		}
	}
	return nil, false
}

func checkAvailability(products []string, commonList *entities.CommonProductList) ([]string, bool) {
	for _, product := range products {
		if (*commonList)[product] <= 0 {
			return nil, false
		}
	}
	return products, true
}

func cook(products []string, commonList *entities.CommonProductList) {
	for _, product := range products {
		(*commonList)[product]--
	}
}
