package main

import "fmt"

func main() {
	fmt.Println("how", "are", "you", "Doing") //... three dots if you see it means you can pass any number of value seperated by comma
	fmt.Println("Sum", sumOf(10, 20, 30, 40, 50, 60, 70))
	fmt.Println("Sum", sumOf(10, 20))
	fmt.Println("Sum", sumOf())
}

// variadic function must be the last parameter
// vraiadic can accept many or even 0 aurgements
// you can pass a slice as a aurgement to a variadic parameter but you need to use ... eclipse sybol append to the slice
func sumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
