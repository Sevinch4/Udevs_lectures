package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8}
	ch := make(chan int, len(a))
	for _, v := range a {
		//v := v
		go func(val int) {
			ch <- val * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
