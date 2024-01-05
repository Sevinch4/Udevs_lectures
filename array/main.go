package main

import "fmt"

func add(n int) {
	slice := make([]int, 0, 10)
	var a int
	f := 0
	for i := 1; i <= n; i++ {
		a = i
		if a%2 == 0 {
			a *= -1
		}
		slice = append(slice, a)
	}
	fmt.Println(slice)
	var x, y int

	for i := 0; i < len(slice); i++ {

		if slice[i] < 0 {
			f++

			if f == 2 {
				x = i
			}
			if f == 4 {
				y = i
				break
			}
		}
	}
	sum := 0

	for i := x + 1; i < y; i++ {
		sum += slice[i]
	}
	fmt.Println(sum)

}

// task-2
func array(arr []int, n int) {
	var (
		nol, musbat, manfiy, juft, toq int = 0, 0, 0, 0, 0
	)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			nol++
		} else if arr[i] > 0 {
			musbat++
		} else if arr[i] < 0 {
			manfiy++
		}
		if arr[i]%2 == 0 {
			juft++
		} else {
			toq++
		}
	}
	fmt.Printf("Musbatlar %dta,Manfiy %dta,Nollar %dta,Juft %d,Toq %d\n", musbat, manfiy, nol, juft, toq)
}

// task-3
func sliceAj(array []int) {
	juft := []int{}
	toq := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			juft = append(juft, array[i])
		} else {
			toq = append(toq, array[i])
		}
	}
	fmt.Println(juft)
	fmt.Println(toq)
}

// task-4
func slice(slice1 []int) {
	juft := 1
	toq := 0
	for i := 0; i < len(slice1); i++ {
		//a := i
		if i%2 == 0 {
			juft *= slice1[i]
		} else {
			toq += slice1[i]
		}
	}
	fmt.Println("Juft (product): ", juft)
	fmt.Println("Toq(sum): ", toq)
}

// taks-5
func middle(slice []int) {
	min := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i+1]
		}
	}
}

//task-6

// task-7
func swap(n int, slice []int) {
	for i := 0; i < n; i++ {
		slice = append(slice, i+1)
	}
	// if i==0 && i==n{
	// 	slice = append(slice, )
	// }

}

// task-8
func max(n int,slice []int) {
	// for i := 0; i < n; i++ {
	// 	slice = append(slice, i+10)
	// }
	max := 0
	a := 0
	for j := 0; j < len(slice); j++ {
		
		if slice[j] > max {
			max = slice[j]
			a = j
		}
	}
	for i := slice[a+1]; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()

}
func main() {
	//task-1
	// var n int
	// fmt.Print("Input num: ")
	// fmt.Scan(&n)
	// add(n)

	//task-2
	// var n1 = 10
	// var arr []int = []int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
	// array(arr, n1)

	//task-3
	// var arr1 []int = []int{2, 3, 5, 4, 12, 78, 90, 45, 76, 99, 87}
	// sliceAj(arr1)

	//task-4
	// slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice(slice1)

	//task-8
	var n2 int = 6
	slice2 := []int {2,5,4,1,3,2}
	max(n2,slice2)
}
