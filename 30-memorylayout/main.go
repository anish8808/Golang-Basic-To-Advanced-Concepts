package main

import (
	"fmt"
	"unsafe"
)

// string size always 16 byte is golang
func main() {
	t1 := T1{ID: 11, name: "Anish", isMarried: false, address: "Bangalore"}
	fmt.Println("t1:", t1, "Size of t1:", unsafe.Sizeof(t1))

	arr1 := []string{"How", "are", "you", "doing"}
	//size := unsafe.Sizeof(arr1[0])

	arrptr := uintptr(unsafe.Pointer(&arr1[0]))
	//fmt.Println(*(*string)(unsafe.Pointer(arrptr)))

	for i := 0; i < 3; i++ {
		arrptr += 16 //string stores 16 byte
		//fmt.Println(*(*string)(unsafe.Pointer(arrptr)))
	}
}

type T1 struct {
	ID        int
	name      string
	isMarried bool
	address   string
}
