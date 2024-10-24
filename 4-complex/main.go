package main

import (
	"fmt"
	"reflect"
)

func main() {
	//float32 , float64 (if dont mention the type the bigger value will be taken by compiler)
	//the below both real and imaginary are float64 and float64
	//c1 is complex128
	c1 := complex(12.13, .5)
	var (
		r1, i1 float32 = 12.31, .5
		r2, i2 float64 = 12.31, .5
	)
	c2 := complex(r1, i1)
	c3 := complex(r2, i2)
	fmt.Println("Value:", c1, "Typw", reflect.TypeOf(c1))
	fmt.Println("Value", c2, "Tye", reflect.TypeOf(c2))
	fmt.Println("Value", c1, "Type", reflect.TypeOf(c3))
	//another declearation
	c4 := 12.31 + .5i
	fmt.Println("Value", c4, "Type", reflect.TypeOf(c4))

	//Operations of complex number
	fmt.Println("Addition of complex numbers:", c1+c3)
	fmt.Println("Substraction of complex numbers:", c1-c3)
	fmt.Println("multiplication of complex numbers:", c1*c3)
	fmt.Println("Divide of complex numbers:", c1/c3)
	//if the type is not same then we can do the type casting
	fmt.Println("Addition of two different type of the complex number ", c1+complex128(c2))

}
