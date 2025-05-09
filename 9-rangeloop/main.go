package main

import "fmt"

// range loop is used in three ways
// 1- arrays , slices --> gives indexx , value
// 2 - maps --> key value
// 3- channels --> value

// range loop is used in three ways
// 1- arrays , slices --> gives indexx , value
// 2 - maps --> key value
// 3- channels --> value
func main() {
	str := "Hello World!"
	for i, v := range str {
		println("index", i, "Value", string(v))
	}
	for _, v := range str { // if you dont want put blank statement
		println("Value", string(v))
	}

	// reverse string logic
	revstr := ""

	for _, v := range str {
		revstr = revstr + string(v)
	}
	println(revstr)

	count := 0
	for range str {
		fmt.Println(str+"->", count)
		count++
	}
}
