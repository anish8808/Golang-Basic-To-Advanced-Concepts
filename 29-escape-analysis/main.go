package main

import (
	"fmt"
	"math/rand"
	"time"
)

// it an algorithm , that decide the particular variable should be in heap
// 3 case
// 1- does not escape -->means store to stack memory
// 2- escape to heap --> complier decide to store in heap
// moved to heap  --> sometimes store in heap (decide by run time )
func main() {
	var num1 int = 100
	const max int = 9999
	var num2 any = num1
	println("num1:", num1, "Max:", max, "num2:", num2)

	slice1 := make([]string, 10000) //--> escape to heap will come
	for i := 0; i <= 10000; i++ {
		slice1 = append(slice1, "")
	}

	fmt.Println(time.Now())
	slice2 := make([]string, 0) //--> does not escape will come
	for i := 0; i <= 1000; i++ {
		slice1 = append(slice2, "Hello World")
	}
	fmt.Println(time.Now()) //println always skip to heap because uses the parameter
	println(slice1)

	sum1 := add(10, 20, 30, 40, 50)
	println("sum1:", sum1)

	sum2 := addPtr(10, 20, 30, 40, 50)
	println("sum1:", *sum2)

	//moved to heap value
	slice4 := generate(100)
	println(slice4)

	//struct

	t1 := new(T1)
	t1.Id = 100
	println(t1.Id)
	t1.slice = []int{10, 20, 30}
	println(t1.slice)

}

func add(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func addPtr(nums ...int) *int { // if we are returning pointer it will be pointer
	sum := new(int)
	for _, v := range nums {
		*sum += v
	}
	return sum
}

//moved to heap examples

func generate(n int) []int {
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10000))
	}
	return arr
}

type T1 struct {
	Id    int
	slice []int
}
