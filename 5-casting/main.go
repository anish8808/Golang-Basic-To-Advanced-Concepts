package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//Casting or conversion
//There is no implicit conversion in go(by default)

func main() {
	var num1 int8 = 100
	//var num2 int16 = num1 NOT OK
	//var num2 int16 = (int16)num1 NOT OK

	var num2 int16 = int16(num1)
	fmt.Println("num1", reflect.TypeOf(num1), "num2", reflect.TypeOf(num2))

	var num3 float32 = 12.14
	var num4 int = int(num3)
	fmt.Println("num2", reflect.TypeOf(num3), "num3", reflect.TypeOf(num4))

	var num5 float32 = 123.456
	var num6 float64 = float64(num5)
	fmt.Println("num5", reflect.TypeOf(num5), "num6", reflect.TypeOf(num6))

	//type conversion

	var num7 uint8 = 65
	var str string = string(num7) //NOT "65" this will print A
	fmt.Println("num7", reflect.TypeOf(num7), "str", reflect.TypeOf(str), "value", str)

	//for "65"
	//this way we can achive this
	str = fmt.Sprint(num7)
	fmt.Println(str)
	//other way we can string to integer
	str1 := strconv.Itoa(int(num7))
	fmt.Println(str1)

	//str2:="a" //convert string to interger NOT possible beacuse string is collection of char
	str2 := 'A' //a character could convert
	fmt.Println(int(str2))
	str3 := "17000"
	num8, _ := strconv.Atoi(str3) //blank identifire (if you are not going to use that variable can use blank identifire)
	fmt.Println(num8)

	//swap
	a, b, c := 10, 20, 30
	fmt.Println("a:", a, "b:", b, "c:", c)
	a, b, c = c, a, b
	fmt.Println("a:", a, "b:", b, "c:", c)
}
