package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Author 1"},
	{ID: 2, Title: "Book 2", Author: "Author 2"},
	{ID: 3, Title: "Book 3", Author: "Author 3"},
}

func main() {
	http.HandleFunc("/books", handleBooks)

	port := 8080
	fmt.Printf("server listening on: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		encodeJson(w, books)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed")
	}
}

func encodeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
