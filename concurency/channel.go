package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	ch := make(chan int)
	////task1
	fmt.Println("------------------------------------------------------------------------------")
	go task1(ch)
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

	//task2
	fmt.Println("------------------------------------------------------------------------------")
	ch2 := make(chan int)
	go task2(ch2)
	fmt.Println("sum: ", <-ch2)

	//task3
	fmt.Println("------------------------------------------------------------------------------")
	ch3 := make(chan int)
	var n int = 3
	go task3(n, ch3)
	fmt.Println("factorial: ", <-ch3)

	//task4
	fmt.Println("------------------------------------------------------------------------------")
	ch4 := make(chan int)
	var n4 int = 91
	go task4(n4, ch4)
	fmt.Println("prime -> ", n4)

	//task5
	fmt.Println("------------------------------------------------------------------------------")
	nums := []int{1, 90, 45, 456, 3, 2, 4, 5, 6}
	ch5 := make(chan int)
	go task5(nums, ch5)
	fmt.Println("max: ", <-ch5)

	//task6
	fmt.Println("------------------------------------------------------------------------------")
	word := "hello"
	ch6 := make(chan []string)
	go task6(word, ch6)
	fmt.Println("reversed: ", <-ch6)

	//task7
	fmt.Println("------------------------------------------------------------------------------")
	slice := []int{4, 3, 2, 6}
	ch7 := make(chan []int)
	go task7(slice, ch7)
	fmt.Println("slice: ", <-ch7)

	//task8
	fmt.Println("------------------------------------------------------------------------------")
	slice8 := []int{2, 4, 6}
	ch8 := make(chan int)
	go task8(slice8, ch8)
	fmt.Println("avg: ", <-ch8)

	//task9
	fmt.Println("------------------------------------------------------------------------------")
	slice9 := []int{1, 3, 6, 123, 13, 8, 45, 23, 56, 123}
	ch9 := make(chan []int)
	go task9(slice9, ch9)
	fmt.Println("sorted: ", <-ch9)

	//task10
	fmt.Println("------------------------------------------------------------------------------")
	ch10 := make(chan int)
	go task10(ch10)
	fmt.Println("sum of square: ", <-ch10)

	//task11
	fmt.Println("------------------------------------------------------------------------------")
	var n1 int = 4
	ch11 := make(chan []int)
	go task11(n1, ch11)
	fmt.Println("fib: ", <-ch11)

	//task12
	fmt.Println("------------------------------------------------------------------------------")
	slice12 := []int{2, 1, 3, 4}
	ch12 := make(chan int)
	go task12(slice12, ch12)
	fmt.Println("sum of even slice: ", <-ch12)

	//task13
	fmt.Println("------------------------------------------------------------------------------")
	var text string = "sevinch university inha "
	ch13 := make(chan int)
	go task13(text, ch13)
	fmt.Println("length: ", <-ch13)

	//task14
	fmt.Println("------------------------------------------------------------------------------")
	fileName := "concurency/channel.go"
	ch14 := make(chan int)
	go task14(fileName, ch14)
	fmt.Println("lines: ", <-ch14)

	//task15
	fmt.Println("------------------------------------------------------------------------------")
	ch15 := make(chan int)
	go producer(ch15)
	go consumer(ch15)

	time.Sleep(time.Millisecond)
}

func task1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Millisecond)
	}
}
func task2(ch chan<- int) {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	ch <- sum
}
func task3(n int, ch chan<- int) {
	fac := 1
	for i := 2; i <= n; i++ {
		fac *= i
	}
	ch <- fac
}
func task4(n int, ch chan<- int) {
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			ch <- i
		}
	}
}
func task5(nums []int, ch chan<- int) {
	max := 0
	for _, m := range nums {
		if m > max {
			max = m
		}
	}
	ch <- max
	time.Sleep(time.Millisecond)
}
func task6(word string, ch chan<- []string) {
	slice := make([]string, 0)
	for i := len(word) - 1; i >= 0; i-- {
		slice = append(slice, string(word[i]))
	}
	ch <- slice
}
func task7(slice []int, ch chan<- []int) {
	slice1 := make([]int, 0)
	//multiple := make(map[int]int)
	for _, s := range slice {
		slice1 = append(slice1, s*2)
	}
	ch <- slice1
}
func task8(slice []int, ch chan<- int) {
	avg := 0
	count := 0
	sum := 0
	for _, s := range slice {
		count++
		sum += s
	}
	avg = sum / count
	ch <- avg
}
func task9(slice []int, ch chan<- []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	ch <- slice
}
func task10(ch chan<- int) {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i * i
	}
	ch <- sum
}
func task11(n int, ch chan<- []int) {
	fib := make([]int, n)

	if n <= 0 {
		return
	}

	fib[0], fib[1] = 0, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	ch <- fib
}
func task12(slice []int, ch chan<- int) {
	sum := 0
	for _, s := range slice {
		if s%2 == 0 {
			sum += s
		}
	}
	ch <- sum
}
func task13(text string, ch chan<- int) {
	words := strings.Split(text, " ")
	max := ""
	for _, w := range words {
		if len(w) > len(max) {
			max = w
		}
	}
	ch <- len(max)
}
func task14(fileName string, ch chan<- int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error is while opening file", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error is while scanning err", err.Error())
		return
	}

	ch <- lineCount
}

// task15
func producer(ch chan<- int) {
	n := rand.Intn(100)
	ch <- n
}
func consumer(ch <-chan int) {
	v := <-ch
	fmt.Println("v: ", v)
}
