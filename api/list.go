package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var n1, n2, n3 string
	fmt.Print("input names: ")
	fmt.Scan(&n1, &n2, &n3)

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, "https://api.agify.io", nil)
	if err != nil {
		panic(err.Error())
		return
	}

	query := request.URL.Query()
	query.Add("name[]", n1)
	query.Add("name[]", n2)
	query.Add("name[]", n3)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
		return
	}

	p := []Person{}

	if err = json.NewDecoder(response.Body).Decode(&p); err != nil {
		panic(err.Error())
		return
	}
	fmt.Println(p)

}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}
