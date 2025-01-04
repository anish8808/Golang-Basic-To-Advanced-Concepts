package main

import "fmt"

func main() {

	var ptr1 *int // what ptr holds is  nil pointer (*ptr1 = 5 is not possible as ptr1 is nil and you are accessing nil run timt error)
	var num1 int = 100

	var ptr2 **int //pointer to pointer

	ptr2 = &ptr1
	ptr1 = &num1
	fmt.Println("Address of num1 is :", &num1)
	fmt.Println("Address of num1 through pointer ptr :", ptr1)
	fmt.Println("value of num1 t :", num1)
	fmt.Println("Value of num1 through pointer ptr :", *ptr1)

	*ptr1 = 200
	fmt.Println("value of num1 t :", num1)
	fmt.Println("Value of num1 through pointer ptr :", *ptr1)

	changeVal(&num1, 500)
	fmt.Println("value of num1 t :", num1)
	fmt.Println("Value of num1 through pointer ptr :", *ptr1)

	changeVal(ptr1, 800)
	fmt.Println("value of num1 t :", num1)
	fmt.Println("Value of num1 through pointer ptr :", *ptr1)

	// if you want to get the num1 value
	fmt.Println("The value of num1 for **ptr2", **ptr2)
	fmt.Println("The address of num1 for **ptr2", *ptr2)

	changeVal(*ptr2, 5001)
	fmt.Println("The value of num1 for **ptr2", **ptr2)
	fmt.Println("The address of num1 for **ptr2", *ptr2)

	//pointer arthmatic is hidden here
	//you can do but not directly

	ptr3 := &num1 //short hand declearation
	fmt.Println("Value:", *ptr3, "Address of num1", ptr3)

	str1 := new(string) // new allocates memeory and return address so str1
	*str1 = "Hello"
	str2 := new(string)
	*str2 = "World"
	fmt.Println(concate(str1, str2))

	str3 := new(string) // type inference here
	str4 := new(string) // it means it will be empty string
	fmt.Println(concate(str3, str4))

}

func changeVal(num *int, val int) {
	*num = val
}

func concate(str1, str2 *string) string {
	return *str1 + *str2
}
