package main

import "fmt"

// import "sort"

// task-1
type Car struct {
	Name       string
	price      int
	model      string
	horsePower int
}

// task-2
type Team struct {
	name        string
	coach       string
	playerCount int
	captain     int
}
type Barcelona struct {
	name string
}

// task-3
type students struct {
	name    string
	age     int
	scholar int
	score   []int
}

func main() {
	//task-1
	// car := []Car{
	// 	{
	// 		Name: "ferrari",
	// 		price: 450,
	// 		model: "ferrari",
	// 		horsePower: 123,
	// 	},
	// 	{
	// 		Name: "merce ",
	// 		price: 10000,
	// 		model: "merce ...",
	// 		horsePower: 1039,
	// 	},
	// 	{
	// 		Name : "bmw",
	// 		price: 900,
	// 		model: "bmw m8",
	// 		horsePower: 700,
	// 	},
	// 	{
	// 		Name : "alp",
	// 		price: 900,
	// 		model: "alp m8",
	// 		horsePower: 700,
	// 	},

	// }
	// for i := 0; i < len(car); i++ {
	// 	for j := i+1; j < len(car); j++ {
	// 		if car[i].model > car[j].model {
	// 			car[j],car[i] = car[i],car[j]
	// 		}
	// 	}
	// }

	// fmt.Println(car)

	//task-2
	//  team := []Team{
	// 	{

	// 	},
	 //}

	//task-3
	students := []students{
		{ // 0
			name:    "sevinch",
			age:     18,
			scholar: 1,
			score:   []int{4, 4, 3, 4, 5},
		},
		{ // 1
			name:    "qobil",
			age:     18,
			scholar: 1,
			score:   []int{4, 5, 4, 5, 4},
		},
		{ // 2
			name:    "nurmuhammad",
			age:     19,
			scholar: 1,
			score:   []int{2, 3, 2, 5, 5},
		},
		{ // 3
			name:    "sarvinoz",
			age:     18,
			scholar: 1,
			score:   []int{3, 4, 5, 2, 5},
		},
	}
	var sum int = 0 // 20
	for i := 0; i < len(students); i++ {
		for j := 0; j < len(students[i].score); j++ {
			sum += students[i].score[j]
		}
		sum /= len(students[i].score)
		if sum == 4 {
			fmt.Println(students[i])
		}
		sum = 0

	}

}
