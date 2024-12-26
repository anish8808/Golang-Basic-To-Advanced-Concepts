package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// slice is expandable array , like can be expend if needed
	// slice is a refrence type

	// slice , map , channel --> to create this we have to make builtin fucn called make

	//nil can be checked on map, pointer , slice , interface
	var slice1 []int // for slice you donnt need to give size

	if slice1 == nil {
		fmt.Println("Slice is nil bracuse it not instantiated")
	}
	slice1 = make([]int, 5) // type inference works here
	fmt.Println(slice1)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(1000)
	}
	fmt.Println(slice1)

	//short hand declearation

	slice2 := []int{10, 20, 30, 40, 50} // slice with value

	fmt.Println(slice2)

	//another way

	slice3 := make([]int, 5, 10) //length =5 , capacity = 10 {it can grow upto 10} if you dont give cap then len will be cap by default
	fmt.Println(slice3, "len", len(slice3), "capacity", cap(slice3))
	//generally cap we need to append till the capacity using append  func same as vec c++

	slice3 = append(slice3, 500) // append the new element
	slice3 = append(slice3, 500) // append the new element
	slice3 = append(slice3, 500) // append the new element
	slice3 = append(slice3, 500) // append the new element
	slice3 = append(slice3, 500) // append the new element
	slice3 = append(slice3, 500) // append the new element
	fmt.Println(slice3, "len", len(slice3), "capacity", cap(slice3))

	slice3 = append(slice3, 505, 506, 507, 508, 509) // it a variadic function means you can add any number of element
	fmt.Println(slice3, "len", len(slice3), "capacity", cap(slice3))
	err := changeVal(slice3, 6, 605)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(slice3)
	}

	//reverse the slice
	reverse(slice3, 0, len(slice3)-1)
	fmt.Println("reversed slice is ", slice3)
}

func changeVal(arr []int, i, v int) error {
	if int(i) > len(arr) {
		return fmt.Errorf("Invalid index")
	}
	arr[i] = v

	return nil
}

//reverse a slice
//deep copy a slice
// by default slice does allow shallow copy beacuse of refernce type

func reverse(arr []int, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}
