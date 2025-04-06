package main

import "fmt"

func main() {
	fmt.Println("Interface based ", sum(10, 10.11, 12, 14, 45, 45.65))
	fmt.Println("Generic based ", sum2[float64](10, 10.11, 12, 14, 45, 45.65))
	fmt.Println("Generic based ", sum2[int](10, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[int](10, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[uint8](10, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[float64](10, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[uint16](10, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[float64](10.0, 10, 12, 14, 45, 45))
	fmt.Println("Generic based ", sumOf[uint16](10.0, 10, 12, 14, 45, 45))
}

func sum1(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

//but we want generic of sum func

func sum(vals ...any) any {
	sum := 0.0

	for _, v := range vals {

		switch v.(type) {
		case int:
			sum = sum + float64(v.(int))
		case float32:
			sum = sum + float64(v.(float32))
		case int64:
			sum = sum + float64(v.(int64))
		case int32:
			sum = sum + float64(v.(int32))
		case int16:
			sum = sum + float64(v.(int16))
		case uint64:
			sum = sum + float64(v.(uint64))
		case float64:
			sum = sum + float64(v.(float64))
		}
	}
	return sum
}

func sum2[T int | float64](vals ...T) T { //generics
	var s T

	for _, v := range vals {
		s += v
	}

	return s
} //generics

type DT interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | float32 | float64
	//this interface is not used for method is should not contain any method
}

func sumOf[T DT](vals ...T) T {

	var s T
	for _, v := range vals {
		s += v
	}

	return s
}
