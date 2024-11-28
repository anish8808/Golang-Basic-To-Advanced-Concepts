package main

import (
	"fmt"
	"reflect"
)

// empty interface --> you can decelare var
func main() {
	//var val interface{}  //this is called interface this is capable of holding any type of data
	var val any //--> any is nothing but empty interface{}

	val = 100
	fmt.Println("Value:", val, "Type:", reflect.TypeOf(val))
	val = false
	fmt.Println(val, reflect.TypeOf(val))
	val = 100.100
	fmt.Println(val, reflect.TypeOf(val))
	val = "Hello World"
	fmt.Println(val, reflect.TypeOf(val))
	val = 'A'
	fmt.Println(val, reflect.TypeOf(val))
	val = byte(65)
	fmt.Println(val, reflect.TypeOf(val))

	//type casting --> v = T(v)
	//type assertion --> = v(T)
	var num1 byte = val.(byte) // its not type casting its type asseretions
	fmt.Println("Value:", num1, "Type:", reflect.TypeOf(num1))

	var a any = 12
	var b int = 12
	var c float32 = 12.00
	println(a.(int) + b + int(c))
	println(float32(a.(int)) + float32(b) + c)
}
