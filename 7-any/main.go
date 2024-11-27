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
}
