package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 10}
	fmt.Println("Add of the 0th element ", &arr1[0])
	fmt.Println("size of int ", unsafe.Sizeof(arr1[0]))
	//reflect.TypeOf()
	size := unsafe.Sizeof(arr1[0])
	arrptr := uintptr(unsafe.Pointer(&arr1[0])) ///this let you perform unsafe operator
	//uintptr is capabel that a store to any type of pointer

	arrptr += size //performing airthmetic operation on uintptr

	val := (*int)(unsafe.Pointer(arrptr)) //this returns a number that will be address so we have to give the start

	fmt.Println(*val)

	arr1ptr := uintptr(unsafe.Pointer(&arr1[0]))
	fmt.Println(*(*int)(unsafe.Pointer(arr1ptr)))
	for i := 0; i < len(arr1); i++ {

		arr1ptr += size
		fmt.Println(*(*int)(unsafe.Pointer(arr1ptr)))
	}
	uintptr1 := uintptr(0x00000000)
	val1 := (*int)(unsafe.Pointer(uintptr1))
	fmt.Println(val1)
}

//go does not support pointer airthmetic
