package main

import (
	calculate "calculator/addition"
	"fmt"
)

func main() {
	performAdd, err := calculate.Add(12, 23.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Perform the addition of two any type of number ", performAdd)
	}

	performMulti, err := calculate.Multiply(13.0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Perform the Multiplicatation of two any type of number ", performMulti)
	}
}
