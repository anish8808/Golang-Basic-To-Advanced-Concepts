package main

import (
	"fmt"
	"reflect"
)

func main() {
	const PT float32 = 3.14 //in go compiler must evalutae the value of a constant at compile time or you can not change the value

	fmt.Printf("value %v , TYPE %T", PT, PT)

	//Group of constat
	//
	const (
		MAX, MIN = 100, 1
	)

	var num = 10

	//const must be evaluated at run time
	const SQUARE1 = 10 * 10
	const SQUARE2 = MAX * MAX

	//this process is not OK beacuse the value will be calculate at run time
	//const SQUARE3 = num*num

	fmt.Println(num)
	fmt.Println("value", SQUARE1, "type", reflect.TypeOf(SQUARE1))
	fmt.Println("value", SQUARE2, "type", reflect.TypeOf(SQUARE2))

}
