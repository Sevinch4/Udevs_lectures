package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func generateRandom() int {
	return rand.Intn(1000)
}
func generate(w http.ResponseWriter, r *http.Request) {
	randomNum := generateRandom()

	str := "random num is : " + strconv.Itoa(randomNum)

	if _, err := w.Write([]byte(str)); err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println("done!")
}

func main() {
	//var name string
	//
	//fmt.Print("enter name: ")
	//fmt.Scan(&name)
	//
	//client := http.Client{}
	//
	//request, err := http.NewRequest(http.MethodGet, "https://api.agify.io", nil)
	//if err != nil {
	//	fmt.Println("error while creating request", err.Error())
	//	return
	//}
	//
	//query := request.URL.Query()
	//query.Add("name", name)
	//request.URL.RawQuery = query.Encode()
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Println("error is while sending request", err.Error())
	//	return
	//}
	//
	////p := []Person{}
	//
	//if err := json.NewDecoder(response.Body).Decode(&p); err != nil {
	//	fmt.Println("error is while decoding body", err.Error())
	//	return
	//}
	//
	//fmt.Println(p)
}

//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Count int    `json:"count"`
//}
