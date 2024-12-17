package main

import "fmt"

func main() {
	num := 80
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough /// by using this it will work similar to c++ and java without break statement in switchcase that make true of all case without checking the case conditions
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough //you can only use fallthrough in switch only
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	default:
		println(num, "is divisible by 2 , 4 , 8")
	}

	var ival interface{}
	var age int16 = 100
	ival = age
	switch ival.(type) {
	case int:
		fmt.Println(ival, "is of type int")
	case int8:
		fmt.Println(ival, "is of type int8")
	case int16:
		fmt.Println(ival, "is of type int16")
	case int32:
		fmt.Println(ival, "is of type int32")
	case int64:
		fmt.Println(ival, "is of type int64")
	default:
		fmt.Println(ival, "is the another type")
	}

}
