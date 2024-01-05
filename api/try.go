package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var name string
	fmt.Print("input name: ")
	fmt.Scan(&name)

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, "https://api.agify.io", nil)
	if err != nil {
		panic(err.Error())
		return
	}

	query := request.URL.Query()
	query.Add("name", name)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		panic(err.Error())
		return
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println("byte is: ", string(b))
}
