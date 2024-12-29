package main

import "fmt"

func main() {
	//1st way declearing
	var arr1 [5]int
	fmt.Println(arr1) // automatically stored with zero of based on data strcture called type inference
	arr1[0] = 100
	arr1[1] = 101
	arr1[2] = 103
	arr1[3] = 104
	arr1[4] = 105

	fmt.Println(arr1)

	sum := 0
	for _, v := range arr1 {
		sum = sum + v
	}
	println("The sum of the all val in the arrays", sum)

	//2nd way declearing
	var arr2 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	//3rd way declearing
	arr3 := [5]int{101, 102, 103, 104, 105}
	fmt.Println(arr3)

	//arr4 the size evaluated by the copiler at compile time
	arr4 := [...]int{123, 234, 2, 4, 434, 34, 33}
	fmt.Println(arr4)

	//legth of the array
	fmt.Println("Length of the array is ", len(arr4))
	fmt.Println("Capacity of the array is ", cap(arr4))

	//loop
	for i := 0; i < len(arr4); i++ {
		print(arr4[i], " ")
	}

	//reverse the array
	start := 0
	end := len(arr4) - 1
	for start < end {
		temp := arr4[start]
		arr4[start] = arr4[end]
		arr4[end] = temp
		end--
		start++
	}
	fmt.Println("Reveresed arr4 is ", arr4)

	//find the minimum in array
	min := arr4[0]
	for i := 0; i < len(arr4); i++ {
		if arr4[i] < min {
			min = arr4[i]
		}
	}
	print("the minimum in array4 is ", min)
}
