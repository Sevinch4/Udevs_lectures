package main

import (
	"fmt"
	"sync"
)

func main() {
	//var startTime = time.Now()
	//fmt.Println(startTime)

	var wg sync.WaitGroup
	wg.Add(2)

	go first(&wg)
	go second(&wg)

	fmt.Println("waiting")
	wg.Wait()
	fmt.Println("done")

	//fmt.Println(time.Since(startTime))
}

func first(wg *sync.WaitGroup) {
	for i := 'A'; i <= 'G'; i++ {
		fmt.Println(string(i))
	}
	wg.Done()
}
func second(wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
