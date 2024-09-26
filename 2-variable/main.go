package main

import (
	"fmt"
	"reflect"
)

func main() {
	//if you use package and varibale and not going to use its an error

	//TYPE-1
	var num1 int

	var char rune //imp  are always goinig to print ascii and after 255 unicode will be print

	char = 19000
	char = 'A'

	fmt.Println("Value and type of num1", num1, reflect.TypeOf(num1))

	fmt.Println("value and type of char ", char, reflect.TypeOf(char))

	//TYPE-2
	var (
		a, b, c int
	)

	fmt.Println(a, b, c)

	//TYPE-3
	var (
		num2, num3 int = 100, 200
	)

	fmt.Println(num2, num3)

	//4th type

	var age = 100 //no need to mention the type her
	fmt.Println(age, reflect.TypeOf(age))

	//5TH TYPE -- shorthand decelare
	num4 := 100 //int by default
	fmt.Printf("Value : %v Type:%T\n", num4, num4)
	num5 := 100.05 // float64
	fmt.Printf("Value : %v Type:%T\n", num5, num5)
	var1 := false // boolean
	fmt.Printf("Value : %v Type:%T\n", var1, var1)
	str1 := "Hello World" //string
	fmt.Printf("Value : %v Type:%T\n", str1, str1)
	str2 := 'A' // rune beacuse single digit
	fmt.Printf("Value : %v Type:%T\n", str2, str2)

}

// NUmbers
// int8 , int16 , int32 , int64
//unit8 , unit16 , uint32 , unit64
//int--> either 4 bytes or 8 bytes depeneding on the architecture
//float --.32 & 64

//bool ---> true/false
//byte--> 1 byte
//rune -- 4 bytes // used to store unicode chars
//string --> collection of chars , Strings are immutable
//complex , complex64 , complex128 --> real and inmaginary value
//builtin  package will be defined here

//type inference
// go does not pick any garbage value
// numbers ->0
//bool --> false
//string ->
