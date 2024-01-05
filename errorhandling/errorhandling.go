package main

import (
	"errors"
	"fmt"
	"log"
	"master/constants"
	"strconv"
)

func main(){

	//error handling - 1
	n := 0
	_,err := fmt.Scan(&n)
	if err != nil{
		fmt.Println("error is ",err)
	}

	//error handling - 2
	nums := []int{12,30,0,90}
	for _,num := range nums{
		result,err := div(num)
		if err != nil{
			if errors.Is(err, constants.ErrSome) {
				fmt.Println("it is true")
			}
			fmt.Printf("Error is %v\n",err)
			return
		}
		fmt.Println(result)
	}

	//atoi
	n1 := "10a"
	s,err := strconv.Atoi(n1)
	if err != nil{
		log.Println("error is: ",err)
		return
	}
	fmt.Printf("%T\n",s)



}
//eror handling - 2
func div(n int)(int,error){
	if n == 0{
		return 0, constants.ErrSome
		//return 0,errors.New("cannot divide by 0")
	}
	return 60/n,nil
}