package main

import "fmt"
func main(){
	n:=1234
	reversed:=0
	for n>0{
		rem:=n%10
		reversed += reversed*10+rem
		n=n/10
	}
	fmt.Println(reversed)
}