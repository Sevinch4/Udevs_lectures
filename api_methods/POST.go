package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = make(map[string]Item)

func main() {
	http.HandleFunc("/items", handleItems)

	fmt.Println("server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		newItem := Item{}
		//var newItem Item
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			fmt.Println(newItem)
			fmt.Print("error:")
			http.Error(w, "error decoding json", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		newItem.ID = fmt.Sprintf("%d", len(items)+1)
		items[newItem.ID] = newItem

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newItem); err != nil {
			fmt.Println("error is whole encodinf json", err.Error())
			return
		}
	default:
		fmt.Println("method not allowed")
		return
	}
}
