package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products = make(map[string]Product)

//var products = map[string]Product{}

func main() {
	products["1"] = Product{
		ID:   "1",
		Name: "olma",
	}
	products["2"] = Product{
		ID:   "2",
		Name: "banan",
	}

	http.HandleFunc("/products/", handleProducts)

	fmt.Println("server running...")
	http.ListenAndServe(":8080", nil)
}
func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		id := r.URL.Path[len("/products/"):]

		if _, exists := products[id]; !exists {
			fmt.Println("error while checking for existence")
			return
		}

		var updatedProduct Product
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			fmt.Println("error is while decoding data", err.Error())
			return
		}
		defer r.Body.Close()

		updatedProduct.ID = id
		products[id] = updatedProduct

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedProduct)
	default:
		fmt.Println("method not allowed")
	}
}
