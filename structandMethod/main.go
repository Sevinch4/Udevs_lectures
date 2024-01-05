package main

import (
	"fmt"
	"master/structandMethod/players"
	"master/structandMethod/workPlace"
)

//embedded struct
type Person struct{
	Name string
	Age int
	Players players.Player
	Workplaces []workplace.Workplace//embedded struct
}
var person = Person{
	Name: "sevinch",
	Age: 18,
	Players: players.Player{
		Name: "Messi",
		Height: 1.60,
	},
	Workplaces: []workplace.Workplace{
		{
			PlaceName: "Udevs",
			Salary: 30_000,
			Level: "Middle",
		},
		{
			PlaceName: "Inha",
			Salary: 20_000,
			Level: "Senior",
		},
	},
}
func main(){
	fmt.Println(person)
	// person.Run()
	// fmt.Println(person.ReturnName())
}