package main

import "fmt"

func main() {
	d := Dog{
		Sound: "vav vav",
	}
	c := Cat{
		Sound: "meow meow",
	}
	h := Horse{
		Sound: "iya iya",
	}

	printSound(d)
	printSound(c)
	printSound(h)





	// s := Square{
	// 	Length: 5,
	// }
	// c := Circle{
	// 	Radius: 4,
	// }
	// t := Triangle{
	// 	Side1: 4,
	// 	Side2: 5,
	// 	Side3: 3,
	// }
	// printShapePerimeter(t)
	// printShapeSurface(t)

	// printShapePerimeter(s)
	// printShapeSurface(s)

	// printShapePerimeter(c)
	// printShapeSurface(c)
}

type Triangle struct {
	Side1 int
	Side2 int 
	Side3 int
}
func (t Triangle) Surface(){
	fmt.Println("Surface : ",(t.Side1*t.Side2*t.Side3))
}
func (t Triangle) Perimeter(){
	fmt.Println("Permieter : ",(t.Side1+t.Side2+t.Side3))
}
type Shape interface {
	Surface()
	Perimeter()
}

func printShapeSurface(s Shape) {
	s.Surface()
}
func printShapePerimeter(s Shape) {
	s.Perimeter()
}

type Square struct {
	Length int
}

func (s Square) Surface() {
	fmt.Println("Surface of square: ", s.Length*s.Length)
}
func (s Square) Perimeter() {
	fmt.Println("Perimeter of square: ", 4*s.Length)
}

type Circle struct {
	Radius int
}

func (c Circle) Surface() {
	fmt.Println("Surface of circle: ", 3*c.Radius)
}
func (c Circle) Perimeter() {
	fmt.Println("Radius of circle: ", 2*3*c.Radius)
}
