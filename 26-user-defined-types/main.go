package main

import (
	"fmt"
	"shapes/shape/cube"
	"shapes/shape/rect"
	"shapes/shape/square"
)

func main() {
	r1 := rect.Rect{
		L: 12.1,
		B: 23.3,
	}

	area := r1.Area()
	permiter := r1.Permiter()
	fmt.Println("Area:", area, "Perimeter:", permiter)
	//or also
	fmt.Println("Area:", r1.A, "Perimeter:", r1.P)

	//2nd way
	c1 := new(cube.Cube)
	c1.L = 10.24
	c1.B = 10.24
	c1.H = 10.24

	fmt.Println("Area:", c1.Area(), "Permetr:", c1.Permiter())

	//3rd way
	s1 := square.Square{
		S: 100,
	}
	fmt.Println("Area:", s1.Area(), "Perimeter:", s1.Permiter())

}
