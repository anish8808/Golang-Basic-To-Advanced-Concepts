package main

import (
	"fmt"
	"math/rand"
)

var (
	a, b, c int // global variables
)

func main() { //{func , if , else , loops , switch you have to put brance imideate to statemnet beacuse go put ; after the each line}
	type char = rune

	type charStr = char

	var age uint8 = 18

	if age >= 18 {
		fmt.Println("age is ", age, " .so eligible for vote")
	} else {
		fmt.Println("age is ", age, "Not eligible for vote")
	}

	var gender rune = 'M' //internally it is 70 bcz actually rune is int32

	if gender >= 21 {
		fmt.Println("age is", age, "and geneder is", string(gender), "eligible for marriage")
	} else if age >= 18 && gender == 'F' {
		fmt.Println("age is", age, "and geneder is", string(gender), "eligible for marriage")
	} else {
		fmt.Println("Not eligible for marriage")
	}

	// conversion either Sprint or strcov

	a, b, c = 72, 49, 64

	if a > b && a > c {
		fmt.Println(a, "a is bigger")
	} else if b > a && b > c {
		fmt.Println(b, "b is bigger")
	} else if c > a && c > b {
		fmt.Println(c, "c is bigger")
	} else if a == b && b == c {
		fmt.Println("equal")
	} else if a == b && b > c {
		fmt.Println("Equal and bigger")
	}

	if num := rand.Intn(1000); num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is a odd number")
	}
}
