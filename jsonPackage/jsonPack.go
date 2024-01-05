package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsStudent bool `json:"isStudent"`
	Courses []string `json:"courses"`
}
func main(){
	//encoding with marshall
	p := Person{
		Name: "Sevinch",
		Age: 18,
		IsStudent: true,
		Courses: []string{"math","history","data structure"},
	}

	jsonData , err := json.Marshal(p)
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println(string(jsonData))
	
	//decoding with unmarshall
	jsonString := `{"name": "Sevinch",
	"age": 18,
	"isStudent": true,
	"courses": ["math","history","data structure"]}`
	var p1 Person
	err1 := json.Unmarshal([]byte(jsonString),&p1)
	if err1 != nil{
		fmt.Println("Eror: ",err)
		return
	}
	fmt.Printf("%+v\n",p1)
}