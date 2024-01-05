package main

import (
	"fmt"
	"main/packages/basket"
	"main/packages/product"
	"main/packages/user"
)
//1:28

func main(){
	products := []product.Product{
		{
			Name: "Suv",
			Price: 2_000,
		},
	}
	b := basket.Basket{
		ProductList: products,
		Total: 2_000,
	}
	u := user.User{
		Name: "Sevinch",
		LastName: "Nabiyeva",
		Basket: b,
	}
	fmt.Println(u)
}