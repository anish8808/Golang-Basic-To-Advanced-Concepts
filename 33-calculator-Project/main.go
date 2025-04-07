package main

import "fmt"

type DT interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | float32 | float64
}

type calculator[T DT] struct {
	A T
	B T
}

func (c calculator[T]) Add() T {
	return c.A + c.B
}

func (c calculator[T]) substract() T {
	return c.A - c.B
}

func (c calculator[T]) multiply() T {
	return c.A * c.B
}

func main() {
	c1 := calculator[int]{A: 10, B: 11}
	c2 := calculator[float32]{A: 10.0, B: 11.0}
	fmt.Println("add c1", c1.Add())
	fmt.Println("add c2", c2.Add())
	fmt.Println("sub c1", c1.substract())
	fmt.Println("sub c2", c2.substract())
	fmt.Println("multilpy c1", c1.multiply())
	fmt.Println("multiply c1", c2.multiply())

}
