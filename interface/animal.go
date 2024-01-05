package main

import "fmt"
 type Animal interface{
	printSound()
 }
 type Dog struct{
	Sound string
 }
 type Cat struct{
	Sound string
 }
 type Horse struct{
	Sound string
 }
 func printSound (a Animal){
	a.printSound()
}

func (d Dog) printSound(){
	fmt.Println("Dog: ",d.Sound)
}
func (c Cat) printSound(){
	fmt.Println("Cat: ",c.Sound)
}
func (h Horse) printSound(){
	fmt.Println("Horse: ",h.Sound)
}