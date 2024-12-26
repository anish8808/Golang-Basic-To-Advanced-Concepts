package main

import (
	"fmt"
	"reflect"
)

func main() {

	//how to create a multi-dimensional array

	arr1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)

	//accessing the element

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			print(arr1[i][j])
		}
		println()
	}

	arr2 := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{1, 2, 3}, {4, 5, 6}}}
	fmt.Println(arr1)

	//accessing the element

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			for k := 0; k < len(arr2[i][j]); k++ {
				print(arr2[i][j][k])
			}
			println()
		}
		println()
	}

	//type on an array

	arr3 := [5]int{49, 34, 54, 67, 89}
	fmt.Println("Values of arr3 is", arr3, "Type of array2", reflect.TypeOf(arr3))

	arr4 := [5]int{}
	arr4 = arr3 //type of the array is same like [5]int for both if this will diff it will not assign
	fmt.Println("Values of arr3 is", arr4, "Type of array2", reflect.TypeOf(arr4))

	err := changeVal(arr3, 1, 334) //call by value {arr will be remian same no impcat}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(arr3)
	}

	max, min, _ := minmax(arr3)
	fmt.Println("max", max, "Min", min)

}

//heap or stack is decided by go run time and compiler not from the dev only we can guess and assume only but not 100 percent sure

//mostly arrays are stored in stack but not guranteed
//escape anylsis (decided by compiler and go rum time memory )

func changeVal(arr [5]int, i int, v int) (err error) {
	if int(i) >= len(arr) {
		//or return errors.New("Invalid Index")
		return fmt.Errorf("Invalid index")
	}
	arr[i] = v
	fmt.Println("inside the function", arr)
	return nil // nil is for interfaces empty interface , pointers , slices , map , channel , maps
}

func minmax(arr [5]int) (max int, min int, err error) { //error as the last we can follow also called idiomatic go
	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("Invalid array")
	}
	min = arr[0]
	max = arr[0]

	for i := 0; i < len(arr); i++ {
		if max > arr[i] {
			max = arr[i]
		}
		if min < arr[i] {
			min = arr[i]
		}
	}
	return max, min, nil
}
