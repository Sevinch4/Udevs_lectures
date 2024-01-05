package main

import (
	"fmt"
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) { //request ob response qaytaradi
	fmt.Println("hello guys")
	_, err := w.Write([]byte(`hello how are you?`)) //to response server
	if err != nil {
		fmt.Errorf("error is while responsing to check endpoint")
		return
	}
	fmt.Println("logic done!")
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("hello / endpoint")
	//})

	//http.HandleFunc("/check", greet)
	//
	//fmt.Println("program is running......")
	//http.ListenAndServe("localhost:8080", nil) //server ishlab turadi

	//request yuvorb response olish
	//var name string
	//fmt.Print("enter name: ")
	//fmt.Scan(&name)
	//
	//client := http.Client{}
	//
	//request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.agify.io?name=%s", name), nil)
	//if err != nil {
	//	fmt.Println("error is while creating request", err.Error())
	//	return
	//}
	//
	//response, err := client.Do(request) //request qberadi va response qaytaradi
	//if err != nil {
	//	fmt.Println("error is while sending request", err.Error())
	//	return
	//}
	//
	//b, err := io.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println("error is while reading body", err.Error())
	//	return
	//}
	//fmt.Println("byte: ", string(b))

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
	fmt.Println("response: ", string(b))
}
