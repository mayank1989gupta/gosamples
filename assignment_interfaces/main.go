package main

import "fmt"

//To create an interface shape and provide print based on shape

//area interface
type shape interface {
	getArea() float64
}

//Struct Triangle
type triangle struct {
	height float64
	base   float64
}

//Struct Square
type square struct {
	sideLength float64
}

func main() {
	s := square{2.4}
	t := triangle{2.4, 3.2}
	//Calling the print method
	printArea(s)
	printArea(t)
}

//Implementing the interface methods for both structs
func printArea(s shape) {
	fmt.Println(s.getArea())
}

//For square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

//For rectangle
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
