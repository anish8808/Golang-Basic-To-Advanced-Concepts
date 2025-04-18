package shape

import (
	"fmt"
)

type IShape interface {
	IArea
	Iperimeter
}

type IArea interface {
	area() float32
}
type Iperimeter interface {
	perimeter() float32
}

// here we are using shape as take interface
type Shape struct {
	ishape IShape
}

//enject ishape in Shape

func New(ishapes IShape) *Shape { //constructor
	return &Shape{ishape: ishapes}
}

func (sh *Shape) Area() {
	fmt.Println("Area of ", sh.ishape.area())
}

func (sh *Shape) Perimeter() {
	fmt.Println("perimeter of ", sh.ishape.perimeter())
}

// insatnces of interface

type rect struct {
	L, B float32
}

func (r *rect) area() float32 {
	return r.L * r.B
}
func (r *rect) perimeter() float32 {
	return 2 * (r.L + r.B)
}

type square float32

func (sq *square) area() float32 {
	return float32(*sq) * float32(*sq)
}
func (sq *square) perimeter() float32 {
	return float32(*sq) + float32(*sq)
}

//finally run fucntion

func RunShape2() {
	r1 := &rect{L: 1.2, B: 2.3}
	r2 := &rect{L: 1.2, B: 2.3}
	r3 := &rect{L: 1.2, B: 2.3}

	s4 := square(2.3)
	s5 := square(1.2)

	shapes := []IShape{r1, r2, r3, &s4, &s5}

	for i, shape := range shapes {
		fmt.Print("shape: ", i+1)
		s := New(shape)
		s.Area()
		s.Perimeter()
		fmt.Println()
	}

}
