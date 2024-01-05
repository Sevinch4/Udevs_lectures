package main

import (
	"fmt"
	"net/http"
)

var users = map[string]string{
	"1": "sarvinoz",
	"2": "sevinch",
	"3": "nurmuhammad",
}

func main() {
	http.HandleFunc("/users/", handleUsers)

	fmt.Println("server running....")
	http.ListenAndServe("localhost:8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		id := r.URL.Path[len("/users/"):]
		if _, exists := users[id]; !exists {
			fmt.Println("error is while checking id")
			return
		}
		delete(users, id)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Item with ID %s deleted\n", id)
	default:
		fmt.Println("method not allowed")
		return
	}
}
