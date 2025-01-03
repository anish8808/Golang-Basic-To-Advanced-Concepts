package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 1, 2, 3, 4, 5}
	slice2 := slice1 // this create a new slice but its a shallow copy(refrence copy)

	slice2[2] = 45

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2) // as slice work on the refernce both will be same shared same memory

	slice3 := slice2[0:6] // 0 to 6th element to slice 3
	fmt.Println("slice3", slice3)
	//slice1 and slice2 and slice 3 everything is mapped with same memort location

	slice3 = append(slice3, 9999)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)

	slice4 := slice1[5:10]
	fmt.Println("slice4", slice4)

	slice5 := slice1[8:] // from 8 to end of the slice
	fmt.Println("slice5", slice5)

	//everyhting in slice is refernce

	//builin function for deepcopy

	slice6 := make([]int, len(slice1))
	copy(slice6, slice1) // deep Copy
	slice6[0] = 11       // now it will not impcat the slice1 as its deepcopy not using same memory
	fmt.Println("slice1", slice1)
	fmt.Println("slice6", slice6)

	//append is a variadic function so 
	slice1 = append(slice1, 10, 20, 30, 40, 50)
	
}
