package user

import "main/packages/basket"

type User struct {
	Name     string
	age      int
	LastName string
	Basket   basket.Basket
}
