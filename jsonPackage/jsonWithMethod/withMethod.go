package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Email string
func (e Email)MarshalJson()([]byte,error){
	return json.Marshal(strings.ToLower(string(e)))
}

func (e *Email) UnmarshalJson(data []byte)error{
	var s string
	err := json.Unmarshal(data,&s)
	if err != nil{
		return err
	}
	*e = Email(strings.ToLower(s))
	return nil
}

type Person struct{
	Name string
	Email Email
}

func main(){
	p := Person{
		Name: "Sevinch Nabiyeva",
		Email: "Sevinch.Nabiyeva.com",
	}

	jsonData,err := json.Marshal(p)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(string(jsonData))

	var decodedPerson Person
	err = json.Unmarshal(jsonData,&decodedPerson)
	if err != nil{
		fmt.Println("error: ",err)
		return 
	}
	fmt.Printf("%+v\n",decodedPerson)
}