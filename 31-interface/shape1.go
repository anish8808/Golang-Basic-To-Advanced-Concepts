package main

import "fmt"

func Area(ishape IShape) {
	fmt.Println("Area", ishape.area())
}
func Perimeter(ishape IShape) {
	fmt.Println("Perimeter", ishape.perimeter())
}

type IShape interface { //interface
	IArea
	IPermiter
}

type IArea interface { // using multiple interface to use spearate or togather as well  // using multiple interface to use spearate or togather as well
	area() float32
}

type IPermiter interface {
	perimeter() float32
}

// creating the new struct and implementing the func area()
type Rect struct {
	L, B float32
}

func (r *Rect) area() float32 {
	return r.L * r.B
}

func (r *Rect) perimeter() float32 {
	return 2 + (r.L + r.B)
}

// //creating the new struct and implementing the func area()
type square float32

func (s *square) area() float32 {
	return float32(*s) * float32(*s)
}

func (s *square) perimeter() float32 {
	return 2 * (float32(*s) + float32(*s))
}

// dependecies enjection or inversion control
// creating a type struct and where Ishape is a intercface
type Shape struct {
	IShape IShape
}

func New(ishape IShape) *Shape { //constructor
	return &Shape{IShape: ishape}
}

func (s *Shape) Area() {
	fmt.Println("Area", s.IShape.area())
}
func (s *Shape) Perimeter() {
	fmt.Println("perimeter ", s.IShape.perimeter())
}

// two or more compnenets can have assosication through interface
// it use mostly in design pattern
// it is a best example of run time binding
func RunShape1() {

	//implementing noramlly
	var check1 = &Rect{L: 1, B: 1}
	check1.area()
	check1.perimeter()

	var check2 = &Rect{L: 1, B: 1}
	check2.area()
	check2.perimeter()

	//impleneting trough interface
	var sq square = 24.3
	s1 := New(&Rect{L: 1.2, B: 2.3}) //this func is working like constructor for Shape struct where Rect will be pass to Interface IShape
	s2 := New(&sq)

	s1.Area()
	s1.Perimeter()
	s2.Area()
	s2.Perimeter()

	//Learning if Rect implements the interface ISShape then rect must implements the func present in the interface
}
