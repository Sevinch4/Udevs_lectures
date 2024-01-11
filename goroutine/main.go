package main

import (
	"fmt"
)

func main() {

	ch := writeWords()
	newCh := removeRepeated(ch)
	// range newCh
	for n := range newCh {
		fmt.Println(n)
	}
	//time.Sleep(time.Second)
}
func writeWords() chan string {
	slice := make([]string, 0)
	slice = []string{"apple", "banana", "banana", "banana", "banana", "banana", "banana", "banana", "apple", "apple", "cherry"}

	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, s := range slice {
			ch <- s
		}
	}()
	return ch
}

func removeRepeated(ch chan string) chan string {
	channel := make(chan string)
	//some := writeWords()
	//v = ch
	go func() {
		defer close(channel)
		v := ""
		for c := range ch {
			if c != v {
				channel <- c
			}
			v = c
		}
	}()
	return channel
}
