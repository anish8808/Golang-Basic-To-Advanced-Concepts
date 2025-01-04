package main

import (
	"errors"
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice1 :", slice1)

	slice1, _ = insert(slice1, 5, 500)
	fmt.Println("slice1 after using insert index:", slice1)
	slice1, _ = insert(slice1, 0, 500)
	fmt.Println("slice1 after using insert index at 0 :", slice1)
	slice2 := []int{100}
	slice2, _ = insert(slice2, 0, 500)
	fmt.Println("slice2 after using insert index:", slice2)

	//deleting the element

	slice1, _ = delete(slice1, 5)
	fmt.Println("slice1 after using delete index:", slice1)

	slice1, _ = delete(slice1, 0)
	fmt.Println("slice1 after using delete index:", slice1)

}
func insert(slice []int, i, v int) ([]int, error) {

	if slice == nil {
		return nil, fmt.Errorf("Invaluid slice ")
	}
	if i < 0 || i > len(slice) {
		return nil, errors.New("Invalid Index")
	}
	slice = append(slice[:i], append([]int{v}, slice[i:]...)...)
	return slice, nil
}

func delete(slices []int, i int) ([]int, error) {
	if slices == nil {
		return nil, fmt.Errorf("Invaluid slice ")
	}
	if i < 0 || i > len(slices) {
		return nil, errors.New("Invalid Index")
	}
	slices = append(slices[:i], slices[i+1:]...)
	return slices, nil
}

